package database

import (
	"strconv"
	"errors"
	"AForum/app/generated/models"
	"AForum/app/generated/restapi/operations/post"
)

// CreateNewPost -
func (m *DB) CreateNewPost(SlugOrID, created string, ps *models.Post) error {
	th, err := m.GetThreadBySlugOrID(SlugOrID)
	if err != nil || th == nil{
		return NotFound(errors.New("Can't find post thread by id: "+SlugOrID))
	}
	frm, fid, err := m.GetForumBySlug(th.Forum)
	if err != nil {
		return NotFound(err)
	}
	usr, uid, err := m.GetUserByName(ps.Author)
	if err != nil {
		return NotFound(errors.New("Can't find post author by nickname: "+ps.Author))
	}
	ps.Author = usr.Nickname
	ps.Thread = th.ID
	ps.Forum = frm.Slug

	{
		if ps.Parent != 0 {
			pr, err := m.GetPostByID(ps.Parent)
			if err != nil {
				return Conflict(errors.New("Parent not found"))
			}
			if pr.Thread != ps.Thread {
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
	`, uid, th.ID, ps.Message, ps.Parent, created).Scan(&ps.Created, &ps.ID); err != nil {
		return AlreadyExist(err)
	}
	if ps.Parent != 0 {
		if _, err := tx.Exec(`
			UPDATE posts
			SET path=(SELECT path FROM posts WHERE pid = $2) || $1::BIGINT
			WHERE pid=$1
		`, ps.ID, ps.Parent); err != nil {
			return err
		}
	} else {
		if _, err := tx.Exec(`
			UPDATE posts
			SET path=ARRAY[$1]
			WHERE pid=$1
		`, ps.ID); err != nil {
			return err
		}
	}
	if _, err := tx.Exec(`
		UPDATE forums 
		SET postCount=postCount+1 
		WHERE fid=$1`, fid); err != nil {
		return err
	}
	return tx.Commit()
}

// GetPosts - 
func (m *DB) GetPosts(params post.ThreadGetPostsParams) (res models.Posts, err error) {
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

	switch *params.Sort {
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
			WHERE p.thread=$1 `+parts["since"]+`
			ORDER BY p.pid `+order+`
			`+parts["limit"]
	case "tree": 
		if params.Since != nil {
			parts["since"] = "AND (path"+ sign +"(SELECT path FROM posts WHERE pid=$"+strconv.Itoa(len(vars)+1)+"))"
			vars = append(vars, params.Since)
		}
		if params.Limit != nil {
			parts["limit"] = "LIMIT $" + strconv.Itoa(len(vars)+1)
			vars = append(vars, params.Limit)
		}
		parts["tail"] = `
			WHERE p.thread=$1 `+parts["since"]+`
			ORDER BY p.path `+order+`
			`+parts["limit"]
	case "parent_tree":
		if params.Since != nil {
			parts["since"] = "AND (pid"+ sign +"(SELECT path[1] FROM posts WHERE pid = $"+strconv.Itoa(len(vars)+1)+"))"
			vars = append(vars, params.Since)
		}
		if params.Limit != nil {
			parts["limit"] = "LIMIT $" + strconv.Itoa(len(vars)+1)
			vars = append(vars, params.Limit)
		}
		parts["tail"] = `
			WHERE p.path[1] IN (
				SELECT pid FROM posts p
				WHERE thread=$1 AND parent=0 `+parts["since"]+`
				ORDER BY p.path[1] `+order+`
				`+parts["limit"]+`
			) ORDER BY p.path[1] `+order+`, p.path ASC`
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
	ps := &models.Post{}
	if err := m.db.QueryRowx(`
		SELECT u.nickname AS author, p.created, f.slug AS forum, p.pid AS id, p.isEdited, p.message, p.parent, t.tid AS thread
		FROM posts   p
		JOIN users   u ON(u.uid=p.author)
		JOIN threads t ON(t.tid=p.thread)
		JOIN forums  f ON(f.fid=t.forum )
		WHERE p.pid=$1
	`, pid).StructScan(ps); err != nil {
		return nil, NotFound(err)
	}
	return ps, nil
}

// GetPost - 
func (m *DB) GetPost(pid int64, related []string) (*models.PostFull, error) {
	ps, err := m.GetPostByID(pid)
	if err != nil {
		return nil, NotFound(err)
	}

	res := &models.PostFull{
		Post: ps,
	}
	for _, r := range related {
		switch r {
		case "user":
			user, _, err := m.GetUserByName(ps.Author)
			if err != nil {
				return nil, err
			}
			res.Author = user
		case "thread":
			th, err := m.GetThreadByID(TID(ps.Thread))
			if err != nil {
				return nil, err
			}
			res.Thread = th

		case "forum":
			frm, _, err := m.GetForumBySlug(ps.Forum)
			if err != nil {
				return nil, err
			}
			res.Forum = frm
		}
	}
	return res, nil
}

// UpdatePost -
func (m *DB)UpdatePost(pid int64, pu *models.PostUpdate) (*models.Post, error) {
	ps, err := m.GetPostByID(pid)
	if err != nil {
		return nil, err
	}
	
	if pu.Message != "" && ps.Message != pu.Message {
		if _, err := m.db.Exec(`
			UPDATE posts
			SET message  = COALESCE(NULLIF($1,''), message),
				isEdited = true
			WHERE pid = $2
		`, pu.Message, pid); err != nil {
			return nil, err
		}
	}
	return m.GetPostByID(pid)
}
