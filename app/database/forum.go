package database

import (
	"github.com/bomkvilt/tech-db-app/app/generated/models"
)

// FID -
type FID int64

// CreateNewForum -
func (m *DB) CreateNewForum(forum *models.Forum) error {
	usr, uid, err := m.GetUserByName(forum.User)
	if err != nil {
		return NotFound(err)
	}
	forum.User = usr.Nickname

	tx := m.db.MustBegin()
	defer tx.Rollback()
	if err := tx.QueryRow(`
		INSERT INTO forums(owner, title, slug, postCount, msgCount, threadCount) 
		VALUES ($1, $2, $3, 0, 0, 0) 
		RETURNING title, slug, postCount, threadCount;
	`, uid, forum.Title, forum.Slug).
		Scan(&forum.Title, &forum.Slug, &forum.Posts, &forum.Threads); err != nil {
		other, _, _ := m.GetForumBySlug(forum.Slug)
		*forum = *other
		return AlreadyExist(err)
	}
	return tx.Commit()
}

func (m *DB) getForum(field string, value interface{}) (frm *models.Forum, fid FID, err error) {
	frm = &models.Forum{}
	if err := m.db.QueryRow(`
		SELECT nickname, title, slug, postCount, threadCount, fid
		FROM forums f
		JOIN users  u ON f.owner=u.uid
		WHERE `+field+`=$1;`, value).
		Scan(&frm.User, &frm.Title, &frm.Slug, &frm.Posts, &frm.Threads, &fid); err != nil {
		return nil, 0, err
	}
	return frm, fid, nil
}

// GetForumBySlug -
func (m *DB) GetForumBySlug(slug string) (*models.Forum, FID, error) {
	return m.getForum("slug", slug)
}

// GetForumByID -
func (m *DB) GetForumByID(fid FID) (*models.Forum, FID, error) {
	return m.getForum("fid", fid)
}
