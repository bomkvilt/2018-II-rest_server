package database

import (
	"AForum/internal/models"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // psql
)

type DB struct {
	db *sqlx.DB
}

// ----------------| common

type conInfo struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

func (ci conInfo) Marshal() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		ci.host, ci.port, ci.user, ci.password, ci.dbName)
}

func New() *DB {
	conn, err := sqlx.Open("postgres", conInfo{
		host:     "127.0.0.1",
		port:     "5432",
		user:     "docker",
		password: "docker",
		dbName:   "docker",
	}.Marshal())
	check(err)

	return &DB{
		db: conn,
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
	if e != nil {
		panic(e)
	}
}

// ----------------|

// @ param string - cmp. sign
// @ param int    - placeholder
type paginatorPart func(string, string) (string, interface{})

type paginator struct {
	order string
	cmp   string
	parts map[string]paginatorPart
	vars  []interface{}
}

func newPaginator() *paginator {
	return &paginator{
		order: "ASC",
		parts: map[string]paginatorPart{},
		vars:  []interface{}{},
	}
}

func (p *paginator) SetOrder(bDesc *bool) *paginator {
	if bDesc != nil && *bDesc {
		p.order = "DESC"
	}
	return p
}

func (p *paginator) GetOrder() string {
	return p.order
}

func (p *paginator) SetCpm(asc, desc string) *paginator {
	if p.order == "asc" {
		p.cmp = asc
	} else {
		p.cmp = desc
	}
	return p
}

func (p *paginator) SetRoot(root interface{}) *paginator {
	p.vars = append(p.vars, root)
	return p
}

func (p *paginator) AddPart(name string, rule paginatorPart) *paginator {
	p.parts[name] = rule
	return p
}

func (p *paginator) AddPartNotNil(name string, rule paginatorPart, target interface{}) *paginator {
	if target != nil {
		return p.AddPart(name, rule)
	}
	return p
}

func (p *paginator) Part(name string) string {
	if _, ok := p.parts[name]; ok {
		str, val := p.parts[name](p.cmp, "$"+strconv.Itoa(len(p.vars)+1))
		if val != nil {
			p.vars = append(p.vars, val)
		}
		return str
	}
	return ""
}

func (p *paginator) Vars() []interface{} {
	return p.vars
}
