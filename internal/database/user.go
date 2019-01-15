package database

import (
	"AForum/internal/models"
	"errors"
	"strings"
)

// InsertNewUser -
func (m *DB) InsertNewUser(u *models.User) error {
	tx, _ := m.db.Begin()
	defer tx.Rollback()

	if _, err := m.db.Exec(`
		INSERT INTO users(nickname, fullname, about, email)
		VALUES ( $1, $2, $3, $4 );
	`, u.Nickname, u.Fullname, u.About, u.Email); err != nil {
		return err
	}
	return tx.Commit()
}

// GetAllCollisions -
func (m *DB) GetAllCollisions(u *models.User) (usrs models.Users, err error) {
	rows, err := m.db.Query(`
		SELECT nickname, fullname, about, email 
		FROM users u
		WHERE u.nickname=$1 OR u.email=$2;
	`, u.Nickname, u.Email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usrs = models.Users{}
	for rows.Next() {
		t := &models.User{}
		if err := rows.Scan(&t.Nickname, &t.Fullname, &t.About, &t.Email); err != nil {
			return nil, err
		}
		usrs = append(usrs, t)
	}
	return usrs, nil
}

func (m *DB) GetForumUsers(params *models.ForumQuery) (usrs models.Users, err error) {
	if !m.CheckForum(params.Slug) {
		return nil, NotFound(errors.New("No forum"))
	}

	q := strings.Builder{}
	q.WriteString(`
		SELECT u.nickname, u.fullname, u.about, u.email
		FROM       forum_users x
		INNER JOIN users       u ON(u.uid=x.username)
		INNER JOIN forums      f ON(f.fid=x.forum)
		WHERE f.slug=$1`)
	if params.Since != "" {
		if params.Desc != nil && *params.Desc {
			q.WriteString(" AND u.nickname < $2")
		} else {
			q.WriteString(" AND u.nickname > $2")
		}
	} else {
		q.WriteString(" AND $2=''")
	}
	if params.Desc != nil && *params.Desc {
		q.WriteString(" ORDER BY u.nickname DESC")
	} else {
		q.WriteString(" ORDER BY u.nickname ASC")
	}
	if params.Limit != nil {
		q.WriteString(" LIMIT $3")
	} else {
		q.WriteString(" LIMIT 99999999+$3")
	}
	rows, err := m.db.Query(q.String(), params.Slug, params.Since, params.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usrs = models.Users{}
	for rows.Next() {
		t := &models.User{}
		rows.Scan(&t.Nickname, &t.Fullname, &t.About, &t.Email)
		usrs = append(usrs, t)
	}
	return usrs, nil
}

// ----------------| Get

func (m *DB) getUser(field string, value interface{}) (u *models.User, err error) {
	u = &models.User{}
	if err := m.db.QueryRow(`
		SELECT nickname, fullname, about, email, uid 
		FROM users 
		WHERE `+field+`=$1;
	`, value).Scan(&u.Nickname, &u.Fullname, &u.About, &u.Email, &u.ID); err != nil {
		return nil, err
	}
	return u, nil
}
func (m *DB) GetUserByName(nick string) (*models.User, error)   { return m.getUser("nickname", nick) }
func (m *DB) GetUserByEmail(email string) (*models.User, error) { return m.getUser("email", email) }
func (m *DB) GetUserByID(uid int64) (*models.User, error)       { return m.getUser("uid", uid) }

func (m *DB) CheckUserByName(nick string) bool {
	dm := 0
	if err := m.db.QueryRow(`
		SELECT uid 
		FROM users 
		WHERE nickname=$1;`, nick).
		Scan(&dm); err != nil {
		return false
	}
	return true
}

func (m *DB) CheckUsersByName(nicks map[string]bool) (map[string]int64, bool) {
	if len(nicks) == 0 {
		return nil, true
	}

	arr := make([]string, 0, len(nicks))
	for n := range nicks {
		arr = append(arr, n)
	}
	rows, err := m.db.Query(`
		SELECT uid
		FROM users 
		WHERE nickname = ANY (ARRAY['` + strings.Join(arr, "', '") + `'])
	`)
	check(err)
	defer rows.Close()

	r := map[string]int64{}
	for _, n := range arr {
		if !rows.Next() {
			return nil, false
		}
		var t int64
		check(rows.Scan(&t))
		r[n] = t
	}
	return r, true
}

// ----------------|

// UpdateUser -
func (m *DB) UpdateUser(u *models.User) error {
	tx, _ := m.db.Begin()
	defer tx.Rollback()

	_, err := tx.Exec(`
		UPDATE users
		SET
			fullname = COALESCE(NULLIF($1, ''), fullname),
			email    = COALESCE(NULLIF($2, ''), email),
			about    = COALESCE(NULLIF($3, ''), about)
		WHERE nickname=$4
		RETURNING fullname, email, about
		`,
		u.Fullname, u.Email, u.About, u.Nickname)
	if err != nil {
		return err
	}
	return tx.Commit()
}
