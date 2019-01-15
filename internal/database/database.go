package database

import (
	"AForum/internal/models"
	"database/sql"

	"github.com/jackc/pgx"
)

type DB struct {
	db *pgx.ConnPool
}

// ----------------| common

func New() *DB {
	db, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "docker",
			Password: "docker",
			Database: "docker",
		},
		MaxConnections: 50,
	})
	check(err)
	return &DB{
		db: db,
	}
}

func (m *DB) Clear() {
	m.db.Exec(`TRUNCATE users, forum_users, forums, threads, votes, posts`)
}

func (m *DB) GetStatus() *models.Status {
	res := &models.Status{}
	m.db.QueryRow(`SELECT count(*) FROM posts`).Scan(&res.Post)
	m.db.QueryRow(`SELECT count(*) FROM threads`).Scan(&res.Thread)
	m.db.QueryRow(`SELECT count(*) FROM forums`).Scan(&res.Forum)
	m.db.QueryRow(`SELECT count(*) FROM users`).Scan(&res.User)
	return res
}

// ----------------| Errors

type (
	ErrorAlreadyExist struct{ error }
	ErrorNotFound     struct{ error }
	ErrorConflict     struct{ error }
)

func (e *ErrorAlreadyExist) Error() string { return e.error.Error() }
func (e *ErrorNotFound) Error() string     { return e.error.Error() }
func (e *ErrorConflict) Error() string     { return e.error.Error() }

func AlreadyExist(e error) error { return &ErrorAlreadyExist{e} }
func NotFound(e error) error     { return &ErrorNotFound{e} }
func Conflict(e error) error     { return &ErrorConflict{e} }

func check(e error) {
	if e != nil && e != sql.ErrNoRows {
		panic(e)
	}
}
