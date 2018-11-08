package database

import (
	"Forum/utiles/walhalla"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	db *sqlx.DB
}

func NewModel(ctx *walhalla.Context) *DB {
	return &DB{
		db: ctx.DB,
	}
}

type (
	ErrorAlreadyExist struct{ error }
	ErrorNotFound     struct{ error }
)

func (e *ErrorAlreadyExist) Error() string { return e.error.Error() }
func (e *ErrorNotFound) Error() string     { return e.error.Error() }

func AlreadyExist(e error) error { return &ErrorAlreadyExist{e} }
func NotFound(e error) error     { return &ErrorNotFound{e} }
