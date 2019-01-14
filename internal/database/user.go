package database

import (
	"AForum/internal/models"
	"strconv"
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
	f, err := m.GetForumBySlug(params.Slug)
	if err != nil {
		return nil, NotFound(err)
	}

	var (
		order = "ASC"
		vars  = make([]interface{}, 1, 3)
		parts = make(map[string]string)
	)
	{ // set flags
		vars[0] = &f.ID
		if params.Desc != nil && *params.Desc {
			order = "DESC"
		}
		if params.Since != nil {
			sign := ">"
			if order == "DESC" {
				sign = "<"
			}

			u, err := m.GetUserByName(*params.Since)
			if err != nil {
				return usrs, nil
			}
			parts["since"] = "AND u.nickname" + sign + "$" + strconv.Itoa(len(vars)+1)
			vars = append(vars, u.Nickname)
		}
		if params.Limit != nil {
			parts["limit"] = "LIMIT $" + strconv.Itoa(len(vars)+1)
			vars = append(vars, params.Limit)
		}
	}

	rows, err := m.db.Query(`
		SELECT u.nickname, u.fullname, u.about, u.email
		FROM       forum_users x
		INNER JOIN users       u ON(u.uid=x.username)
		WHERE x.forum=$1    `+parts["since"]+`
		ORDER BY u.nickname `+order+`
		`+parts["limit"]+`
	`, vars...)
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
