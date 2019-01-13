package database

import (
	"errors"
	"strconv"

	"AForum/internal/models"
)

// CreateNewPost -
func (m *DB) CreateNewPost(SlugOrID, created string, p *models.Post) error {
	t, err := m.GetThreadBySlugOrID(SlugOrID)
	if err != nil {
		return NotFound(errors.New("Can't find post thread by id: " + SlugOrID))
	}
	f, err := m.GetForumBySlug(t.Forum)
	if err != nil {
		return NotFound(err)
	}
	u, err := m.GetUserByName(p.Author)
	if err != nil {
		return NotFound(errors.New("Can't find post author by nickname: " + p.Author))
	}
	p.Author = u.Nickname
	p.Thread = t.ID
	p.Forum = f.Slug

	{
		if p.Parent != 0 {
			pr, err := m.GetPostByID(p.Parent)
			if err != nil {
				return Conflict(errors.New("Parent not found"))
			}
			if pr.Thread != p.Thread {
				return Conflict(errors.New("Parent post was created in another thread"))
			}
		}
	}

	tx := m.db.MustBegin()
	defer tx.Rollback()
	if err := tx.QueryRow(`
		INSERT INTO posts(author, thread, message, parent, isEdited, created, path)
		VALUES ($1, $2, $3, $4, false, $5, ARRAY[]::BIGINT[])
		RETURNING created, pid
	`, u.ID, t.ID, p.Message, p.Parent, created).Scan(&p.Created, &p.ID); err != nil {
		return AlreadyExist(err)
	}
	if p.Parent != 0 {
		if _, err := tx.Exec(`
			UPDATE posts
			SET path=(SELECT path FROM posts WHERE pid = $2) || $1::BIGINT
			WHERE pid=$1
		`, p.ID, p.Parent); err != nil {
			return err
		}
	} else {
		if _, err := tx.Exec(`
			UPDATE posts
			SET path=ARRAY[$1]
			WHERE pid=$1
		`, p.ID); err != nil {
			return err
		}
	}
	if _, err := tx.Exec(`
		UPDATE forums 
		SET postCount=postCount+1 
		WHERE fid=$1`, f.ID); err != nil {
		return err
	}
	return tx.Commit()
}

// GetPosts -
func (m *DB) GetPosts(params *models.PostQuery) (res models.Posts, err error) {
	th, err := m.GetThreadBySlugOrID(params.SlugOrID)
	if err != nil {
		return nil, NotFound(err)
	}

	var (
		sign  = ">"
		order = "ASC"
		vars  = make([]interface{}, 1)
		parts = map[string]string{}
	)
	if params.Desc != nil && *params.Desc {
		order = "DESC"
		sign = "<"
	}
	vars[0] = &th.ID

	switch params.Sort {
	case "flat":
		if params.Since != nil {
			parts["since"] = "AND p.pid" + sign + "$" + strconv.Itoa(len(vars)+1)
			vars = append(vars, params.Since)
		}
		if params.Limit != nil {
			parts["limit"] = "LIMIT $" + strconv.Itoa(len(vars)+1)
			vars = append(vars, params.Limit)
		}
		parts["tail"] = `
			WHERE p.thread=$1 ` + parts["since"] + `
			ORDER BY p.pid ` + order + `
			` + parts["limit"]
	case "tree":
		if params.Since != nil {
			parts["since"] = "AND (path" + sign + "(SELECT path FROM posts WHERE pid=$" + strconv.Itoa(len(vars)+1) + "))"
			vars = append(vars, params.Since)
		}
		if params.Limit != nil {
			parts["limit"] = "LIMIT $" + strconv.Itoa(len(vars)+1)
			vars = append(vars, params.Limit)
		}
		parts["tail"] = `
			WHERE p.thread=$1 ` + parts["since"] + `
			ORDER BY p.path ` + order + `
			` + parts["limit"]
	case "parent_tree":
		if params.Since != nil {
			parts["since"] = "AND (pid" + sign + "(SELECT path[1] FROM posts WHERE pid = $" + strconv.Itoa(len(vars)+1) + "))"
			vars = append(vars, params.Since)
		}
		if params.Limit != nil {
			parts["limit"] = "LIMIT $" + strconv.Itoa(len(vars)+1)
			vars = append(vars, params.Limit)
		}
		parts["tail"] = `
			WHERE p.path[1] IN (
				SELECT pid FROM posts p
				WHERE thread=$1 AND parent=0 ` + parts["since"] + `
				ORDER BY p.path[1] ` + order + `
				` + parts["limit"] + `
			) ORDER BY p.path[1] ` + order + `, p.path ASC`
	}
	rows, err := m.db.Query(`
		SELECT u.nickname, p.created, f.slug, p.pid, p.isEdited, p.message, p.parent
		FROM posts   p
		JOIN users   u ON p.author=u.uid
		JOIN threads t ON p.thread=t.tid
		JOIN forums  f ON t.forum=f.fid
		`+parts["tail"]+`
	`, vars...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res = models.Posts{}
	for rows.Next() {
		tmp := &models.Post{
			Thread: th.ID,
		}
		if err := rows.Scan(&tmp.Author, &tmp.Created, &tmp.Forum, &tmp.ID, &tmp.IsEdited, &tmp.Message, &tmp.Parent); err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}
	return res, nil
}

// GetPostByID -
func (m *DB) GetPostByID(pid int64) (*models.Post, error) {
	p := &models.Post{}
	if err := m.db.QueryRowx(`
		SELECT u.nickname AS author, p.created, f.slug AS forum, p.pid AS id, p.isEdited, p.message, p.parent, t.tid AS thread
		FROM posts   p
		JOIN users   u ON(u.uid=p.author)
		JOIN threads t ON(t.tid=p.thread)
		JOIN forums  f ON(f.fid=t.forum )
		WHERE p.pid=$1
	`, pid).StructScan(p); err != nil {
		return nil, NotFound(err)
	}
	return p, nil
}

// GetPost -
func (m *DB) GetPost(pid int64, related []string) (*models.PostFull, error) {
	p, err := m.GetPostByID(pid)
	if err != nil {
		return nil, NotFound(err)
	}

	res := &models.PostFull{
		Post: p,
	}
	for _, r := range related {
		switch r {
		case "user":
			u, err := m.GetUserByName(p.Author)
			if err != nil {
				return nil, err
			}
			res.Author = u
		case "thread":
			th, err := m.GetThreadByID(p.Thread)
			if err != nil {
				return nil, err
			}
			res.Thread = th

		case "forum":
			f, err := m.GetForumBySlug(p.Forum)
			if err != nil {
				return nil, err
			}
			res.Forum = f
		}
	}
	return res, nil
}

// UpdatePost -
func (m *DB) UpdatePost(p *models.Post) error {
	cp, err := m.GetPostByID(p.ID)
	if err != nil {
		return err
	}

	if p.Message != "" && cp.Message != p.Message {
		if _, err := m.db.Exec(`
			UPDATE posts
			SET message  = COALESCE(NULLIF($1,''), message),
				isEdited = true
			WHERE pid = $2
		`, p.Message, p.ID); err != nil {
			return err
		}
	}
	o, err := m.GetPostByID(p.ID)
	if err == nil {
		*p = *o
	}
	return err
}
