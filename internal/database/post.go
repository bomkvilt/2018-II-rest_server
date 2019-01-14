package database

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lib/pq"

	"AForum/internal/models"
)

func (m *DB) CreateNewPosts(SlugOrID, created string, ps models.Posts) error {
	t, err := m.GetThreadBySlugOrID(SlugOrID)
	if err != nil {
		return NotFound(errors.New("Can't find post thread by id: " + SlugOrID))
	}
	f, err := m.GetForumBySlug(t.Forum)
	if err != nil {
		return NotFound(err)
	}

	if len(ps) == 0 {
		return nil
	}

	parentsSet := map[int64]bool{}
	authorsSet := map[string]bool{}
	for _, p := range ps {
		if p.Parent != 0 {
			parentsSet[p.Parent] = true
		}
		authorsSet[p.Author] = true
	}

	parents := map[int64]*models.Post{}
	for pid := range parentsSet {
		pr, err := m.GetPostByID(pid)
		if err != nil {
			return Conflict(errors.New("Parent not found " + strconv.FormatInt(pid, 10)))
		}
		if pr.Thread != t.ID {
			return Conflict(errors.New("Parent post was created in another thread"))
		}
		parents[pid] = pr
	}
	if !m.CheckUsersByName(authorsSet) {
		return NotFound(errors.New("Can't find post author"))
	}

	for _, p := range ps {
		p.Thread = t.ID
		p.Forum = f.Slug
		if r := parents[p.Parent]; r != nil {
			p.Path = r.Path
		} else {
			p.Path = []int64{}
		}
	}
	b, a := m.postQueryBuilder(ps, t.ID, created)

	//
	tx := m.db.MustBegin()
	defer tx.Rollback()

	rows, err := m.db.Query(b.String(), a...)
	if err != nil {
		return Conflict(err)
	}
	defer rows.Close()

	for _, p := range ps {
		rows.Next()
		if err := rows.Scan(&p.Created, &p.ID); err != nil {
			return Conflict(err)
		}
	}
	return tx.Commit()
}

func (m *DB) postQueryBuilder(ps models.Posts, tid int, created string) (*strings.Builder, []interface{}) {
	b := &strings.Builder{}
	a := []interface{}{}

	b.WriteString("INSERT INTO posts(author, thread, message, parent, isEdited, created, path, forum) VALUES ")
	for i, p := range ps {
		if i != 0 {
			b.WriteString(", ")
		}

		c := 7 * i
		b.WriteString(fmt.Sprintf("($%d, $%d, $%d, $%d, false, $%d, $%d, $%d)",
			c+1, c+2, c+3, c+4, c+5, c+6, c+7))
		a = append(a, p.Author, tid, p.Message, p.Parent, created, pq.Array(p.Path), p.Forum)
	}
	b.WriteString("RETURNING created, pid")

	return b, a
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
			parts["since"] = "AND pid" + sign + "$" + strconv.Itoa(len(vars)+1)
			vars = append(vars, params.Since)
		}
		if params.Limit != nil {
			parts["limit"] = "LIMIT $" + strconv.Itoa(len(vars)+1)
			vars = append(vars, params.Limit)
		}
		parts["tail"] = `
			WHERE thread=$1 ` + parts["since"] + `
			ORDER BY pid    ` + order + `
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
			WHERE thread=$1 ` + parts["since"] + `
			ORDER BY path   ` + order + `
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
			WHERE path[1] IN (
				SELECT pid FROM posts
				WHERE thread=$1 AND parent=0 ` + parts["since"] + `
				ORDER BY path[1] ` + order + `
				` + parts["limit"] + `
			) ORDER BY path[1] ` + order + `, path ASC`
	}
	rows, err := m.db.Query(`
		SELECT author, created, forum, pid, isEdited, message, parent, thread, path
		FROM posts
		`+parts["tail"]+`
	`, vars...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res = models.Posts{}
	for rows.Next() {
		t := &models.Post{
			Thread: th.ID,
		}
		if err := rows.Scan(&t.Author, &t.Created, &t.Forum, &t.ID, &t.IsEdited, &t.Message, &t.Parent, &t.Thread, pq.Array(&t.Path)); err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

// GetPostByID -
func (m *DB) GetPostByID(pid int64) (*models.Post, error) {
	p := &models.Post{}
	if err := m.db.QueryRow(`
		SELECT   author   , created   , forum   , pid  , isEdited   , message   , parent   , thread   , path 
		FROM posts 
		WHERE pid=$1
	`, pid).Scan(&p.Author, &p.Created, &p.Forum, &p.ID, &p.IsEdited, &p.Message, &p.Parent, &p.Thread, pq.Array(&p.Path)); err != nil {
		println(err.Error(), "GetPostByID")
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
	if p.Message != "" {
		if _, err := m.db.Exec(`
			UPDATE posts
			SET message = $1
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
