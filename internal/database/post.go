package database

import (
	// "time"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx"
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
	uids, ok := m.CheckUsersByName(authorsSet)
	if !ok {
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
	tx, _ := m.db.Begin()
	defer tx.Rollback()

	rows, err := tx.Query(b.String(), a...)
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
	rows.Close()

	if _, err := tx.Exec(`
		UPDATE forums 
		SET postCount=postCount+$2
		WHERE fid=$1
	`, f.ID, len(ps)); err != nil {
		return Conflict(err)
	}

	c, d := m.postForumUsersQueryBuilder(ps, f.ID, uids)
	if _, err := tx.Exec(c.String(), d...); err != nil {
		return Conflict(err)
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
	b.WriteString(" RETURNING created, pid")

	return b, a
}

func (m *DB) postForumUsersQueryBuilder(ps models.Posts, fid int, uids map[string]int64) (*strings.Builder, []interface{}) {
	b := &strings.Builder{}
	a := []interface{}{}

	b.WriteString("INSERT INTO forum_users (forum, username) VALUES ")
	for i, p := range ps {
		if i != 0 {
			b.WriteString(", ")
		}

		c := 2 * i
		b.WriteString(fmt.Sprintf("($%d, $%d)", c+1, c+2))
		a = append(a, fid, uids[p.Author])
	}
	b.WriteString(" ON conflict do nothing")

	return b, a
}

// GetPosts -
func (m *DB) GetPosts(params *models.PostQuery) (res models.Posts, err error) {
	th, err := m.GetThreadBySlugOrID(params.SlugOrID)
	if err != nil {
		return nil, NotFound(err)
	}

	q := strings.Builder{}
	switch params.Sort {
	case "flat":
		q.WriteString(`
			SELECT author, created, forum, pid, isEdited, message, parent, thread, path
			FROM posts
			WHERE thread = $1`)
		if params.Since != 0 {
			if params.Desc != nil && *params.Desc {
				q.WriteString(" AND pid < $2")
			} else {
				q.WriteString(" AND pid > $2")
			}
		} else {
			q.WriteString(" AND $2 = 0")
		}
		if params.Desc != nil && *params.Desc {
			q.WriteString(" ORDER BY pid DESC")
		} else {
			q.WriteString(" ORDER BY pid ASC")
		}
		q.WriteString(" LIMIT $3")
	case "tree":
		q.WriteString(`
			SELECT author, created, forum, pid, isEdited, message, parent, thread, path
			FROM posts
			WHERE thread = $1`)
		if params.Since != 0 {
			if params.Desc != nil && *params.Desc {
				q.WriteString(" AND path < (SELECT path FROM posts WHERE pid = $2)")
			} else {
				q.WriteString(" AND path > (SELECT path FROM posts WHERE pid = $2)")
			}
		} else {
			q.WriteString(" AND $2 = 0")
		}
		if params.Desc != nil && *params.Desc {
			q.WriteString(" ORDER BY path DESC")
		} else {
			q.WriteString(" ORDER BY path ASC")
		}
		q.WriteString(" LIMIT $3")
	case "parent_tree":
		q.WriteString(`
			SELECT string_agg(pid::text, ', ')
			FROM (
				SELECT pid FROM posts
				WHERE thread=$1 AND parent=0`)
		if params.Since != 0 {
			if params.Desc != nil && *params.Desc {
				q.WriteString(" AND pid < (SELECT path[1] FROM posts WHERE pid = $2)")
			} else {
				q.WriteString(" AND pid > (SELECT path[1] FROM posts WHERE pid = $2)")
			}
		} else {
			q.WriteString(" AND $2 = 0")
		}
		if params.Desc != nil && *params.Desc {
			q.WriteString(" ORDER BY pid DESC")
		} else {
			q.WriteString(" ORDER BY pid ASC")
		}
		q.WriteString(" LIMIT $3) x")
	}

	// t := time.Now()

	tx, _ := m.db.Begin()
	defer tx.Rollback()

	// t1 := time.Since(t) / time.Microsecond
	
	var rows *pgx.Rows
	if params.Sort == "parent_tree" {
		var points string
		tx.QueryRow(q.String(), th.ID, params.Since, params.Limit).Scan(&points)

		q.Reset()
		q.WriteString(`
			SELECT author, created, forum, pid, isEdited, message, parent, thread, path
			FROM posts
			WHERE path[1] = ANY ('{ `+points+` }'::BIGINT[])`)
		if params.Desc != nil && *params.Desc {
			q.WriteString(" ORDER BY path[1] DESC, path")
		} else {
			q.WriteString(" ORDER BY path[1] ASC, path")
		}
		rows, err = tx.Query(q.String())
	} else {
		rows, err = tx.Query(q.String(), th.ID, params.Since, params.Limit)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// t2 := time.Since(t) / time.Microsecond

	res = make(models.Posts, 0, 20)
	for rows.Next() {
		t := &models.Post{}
		t.Thread = th.ID
		rows.Scan(&t.Author, &t.Created, &t.Forum, &t.ID, &t.IsEdited, &t.Message, &t.Parent, &t.Thread, pq.Array(&t.Path))
		res = append(res, t)
	}

	// t3 := time.Since(t) / time.Microsecond
	// println(t1, t2, t3, "------------")

	return res, tx.Commit()
}

// GetPostByID -
func (m *DB) GetPostByID(pid int64) (*models.Post, error) {
	p := &models.Post{}
	if err := m.db.QueryRow(`
		SELECT   author   , created   , forum   , pid  , isEdited   , message   , parent   , thread   , path 
		FROM posts 
		WHERE pid=$1
	`, pid).Scan(&p.Author, &p.Created, &p.Forum, &p.ID, &p.IsEdited, &p.Message, &p.Parent, &p.Thread, pq.Array(&p.Path)); err != nil {
		// println(err.Error(), "GetPostByID")
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
