package database

import (
	"AForum/internal/models"
	"strconv"
)

// InsertNewUser -
func (db *DB) InsertNewUser(u *models.User) error {
	tx := db.db.MustBegin()
	defer tx.Rollback()

	stmt, err := tx.Preparex(`
	INSERT INTO users(nickname, fullname, about, email)
	VALUES ( $1, $2, $3, $4 );
	`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(u.Nickname, u.Fullname, u.About, u.Email); err != nil {
		return err
	}
	return tx.Commit()
}

// GetAllCollisions -
func (db *DB) GetAllCollisions(u *models.User) (usrs models.Users, err error) {
	rows, err := db.db.Query(`
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
		tmp := &models.User{}
		err := rows.Scan(&tmp.Nickname, &tmp.Fullname, &tmp.About, &tmp.Email)
		if err != nil {
			return nil, err
		}
		usrs = append(usrs, tmp)
	}
	return usrs, nil
}

func (db *DB) GetForumUsers(params *models.ForumQuery) (usrs models.Users, err error) {
	f, err := db.GetForumBySlug(params.Slug)
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

			u, err := db.GetUserByName(*params.Since)
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

	rows, err := db.db.Query(`
		SELECT DISTINCT u.nickname, u.fullname, u.about, u.email
		FROM      users   u
		LEFT JOIN posts   p  ON(u.uid=p.author )
		LEFT JOIN threads t  ON(t.tid=p.thread )
		LEFT JOIN threads t2 ON(u.uid=t2.author)
		WHERE (t.forum=$1 OR t2.forum=$1) `+parts["since"]+`
		ORDER BY u.nickname `+order+`
		`+parts["limit"]+`
	`, vars...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usrs = models.Users{}
	for rows.Next() {
		tmp := &models.User{}
		err := rows.Scan(&tmp.Nickname, &tmp.Fullname, &tmp.About, &tmp.Email)
		if err != nil {
			return nil, err
		}
		usrs = append(usrs, tmp)
	}
	return usrs, nil
}

// ----------------| Get

func (db *DB) getUser(field string, value interface{}) (u *models.User, err error) {
	u = &models.User{}
	if err := db.db.QueryRow(`
		SELECT nickname, fullname, about, email, uid 
		FROM users 
		WHERE `+field+`=$1;`, value).
		Scan(&u.Nickname, &u.Fullname, &u.About, &u.Email, &u.ID); err != nil {
		return nil, err
	}
	return u, nil
}
func (db *DB) GetUserByName(nick string) (*models.User, error)   { return db.getUser("nickname", nick) }
func (db *DB) GetUserByEmail(email string) (*models.User, error) { return db.getUser("email", email) }
func (db *DB) GetUserByID(uid int64) (*models.User, error)       { return db.getUser("uid", uid) }

// ----------------|

// UpdateUser -
func (db *DB) UpdateUser(u *models.User) error {
	tx := db.db.MustBegin()
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
