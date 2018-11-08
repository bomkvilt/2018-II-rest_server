package database

import (
	"Forum/app/generated/models"
)

func (m *DB) CreateNewPost(SlugOrID, created string, ps *models.Post) error {
	th, err := m.GetThreadBySlugOrID(SlugOrID)
	if err != nil {
		return NotFound(err)
	}
	frm, _, err := m.GetForumBySlug(th.Forum)
	if err != nil {
		return NotFound(err)
	}
	usr, uid, err := m.GetUserByName(ps.Author)
	if err != nil {
		return NotFound(err)
	}
	ps.Author = usr.Nickname
	ps.Thread = th.ID
	ps.Forum = frm.Slug

	tx := m.db.MustBegin()
	defer tx.Rollback()
	if err := tx.QueryRow(`
		INSERT INTO posts(author, thread, message, parent, isEdited, created)
		VALUES ($1, $2, $3, 0, false, $4)
		RETURNING created, pid
	`, uid, th.ID, ps.Message, created).Scan(&ps.Created, &ps.ID); err != nil {
		return AlreadyExist(err)
	}
	return tx.Commit()
}
