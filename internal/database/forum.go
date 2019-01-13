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

	tx := m.db.MustBegin()
	defer tx.Rollback()
	if err := tx.QueryRow(`
		INSERT INTO forums(owner, title, slug, postCount, msgCount, threadCount) 
		VALUES ($1, $2, $3, 0, 0, 0) 
		RETURNING title, slug, postCount, threadCount;
		`, u.ID, f.Title, f.Slug).
		Scan(&f.Title, &f.Slug, &f.Posts, &f.Threads); err != nil {
		o, _ := m.GetForumBySlug(f.Slug)
		*f = *o
		return AlreadyExist(err)
	}
	return tx.Commit()
}

func (m *DB) getForum(field string, value interface{}) (f *models.Forum, err error) {
	f = &models.Forum{}
	if err := m.db.QueryRow(`
		SELECT nickname, title, slug, postCount, threadCount, fid
		FROM forums f
		JOIN users  u ON f.owner=u.uid
		WHERE `+field+`=$1;`, value).
		Scan(&f.Author, &f.Title, &f.Slug, &f.Posts, &f.Threads, &f.ID); err != nil {
		return nil, err
	}
	return f, nil
}
func (m *DB) GetForumBySlug(slug string) (*models.Forum, error) { return m.getForum("slug", slug) }
func (m *DB) GetForumByID(fid int64) (*models.Forum, error)     { return m.getForum("fid", fid) }
