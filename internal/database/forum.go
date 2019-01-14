package database

import (
	"AForum/internal/models"
)

// CreateNewForum -
func (m *DB) CreateNewForum(f *models.Forum) error {
	u, err := m.GetUserByName(f.Author)
	if err != nil {
		return NotFound(err)
	}
	f.Author = u.Nickname

	tx, _ := m.db.Begin()
	defer tx.Rollback()
	if err := tx.QueryRow(`
		INSERT INTO forums(owner, title, slug, threadCount, postCount) 
		VALUES ($1, $2, $3, 0, 0) 
		RETURNING fid; 
	`, u.ID, f.Title, f.Slug).Scan(&f.ID); err != nil {
		o, e := m.GetForumBySlug(f.Slug)
		check(e)
		*f = *o
		return AlreadyExist(err)
	}
	return tx.Commit()
}

func (m *DB) getForum(field string, value interface{}) (f *models.Forum, err error) {
	f = &models.Forum{}
	if err := m.db.QueryRow(`
		SELECT nickname, title, slug, threadCount, postCount, fid
		FROM forums f
		JOIN users  u ON f.owner=u.uid
		WHERE `+field+`=$1;
	`, value).Scan(&f.Author, &f.Title, &f.Slug, &f.Threads, &f.Posts, &f.ID); err != nil {
		return nil, err
	}
	return f, nil
}
func (m *DB) GetForumBySlug(slug string) (*models.Forum, error) { return m.getForum("slug", slug) }
func (m *DB) GetForumByID(fid int64) (*models.Forum, error)     { return m.getForum("fid", fid) }

func (m *DB) CheckForum(slug string) bool {
	var tester int
	return m.db.QueryRow("SELECT fid FROM forums WHERE slug=$1", slug).Scan(&tester) == nil
}
