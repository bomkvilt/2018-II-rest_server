package database

import (
	"github.com/bomkvilt/tech-db-ap/app/generated/models"
	"github.com/bomkvilt/tech-db-ap/app/generated/restapi/operations/forum"
	"strconv"
)

// UID type
type UID int

// InsertNewUser -
func (m *DB) InsertNewUser(usr *models.User) error {
	tx := m.db.MustBegin()
	defer tx.Rollback()

	stmt, err := tx.Preparex(`
	INSERT INTO users(nickname, fullname, about, email)
	VALUES ( $1, $2, $3, $4 );
	`)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(usr.Nickname, usr.Fullname, usr.About, usr.Email); err != nil {
		return err
	}
	return tx.Commit()
}

// GetAllCollisions -
func (m *DB) GetAllCollisions(usr *models.User) (usrs models.Users, err error) {
	rows, err := m.db.Query(`
	SELECT nickname, fullname, about, email 
	FROM users u
	WHERE u.nickname=$1 OR u.email=$2;
	`, usr.Nickname, usr.Email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

func (m *DB) GetForumUsers(params forum.ForumGetUsersParams) (usrs models.Users, err error) {
	_, fid, err := m.GetForumBySlug(params.Slug)
	if err != nil {
		return nil, NotFound(err)
	}

	var (
		order = "ASC"
		vars  = make([]interface{}, 1, 3)
		parts = make(map[string]string)
	)
	vars[0] = &fid
	if params.Desc != nil && *params.Desc {
		order = "DESC"
	}
	if params.Since != nil {
		sign := ">"
		if order == "DESC" {
			sign = "<"
		}

		usr, _, err := m.GetUserByName(*params.Since)
		if err != nil {
			return usrs, nil
		}
		parts["since"] = "AND u.nickname" + sign + "$" + strconv.Itoa(len(vars)+1)
		vars = append(vars, usr.Nickname)
	}
	if params.Limit != nil {
		parts["limit"] = "LIMIT $" + strconv.Itoa(len(vars)+1)
		vars = append(vars, params.Limit)
	}

	rows, err := m.db.Query(`
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
		return usrs, nil
	}
	defer rows.Close()

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

func (m *DB) getUser(field string, value interface{}) (usr *models.User, uid UID, err error) {
	usr = &models.User{}
	if err := m.db.QueryRow(`
		SELECT nickname, fullname, about, email, uid 
		FROM users 
		WHERE `+field+`=$1;`, value).
		Scan(&usr.Nickname, &usr.Fullname, &usr.About, &usr.Email, &uid); err != nil {
		return nil, 0, err
	}
	return usr, uid, nil
}

// GetUserByName -
func (m *DB) GetUserByName(nickname string) (*models.User, UID, error) {
	return m.getUser("nickname", nickname)
}

// GetUserByEmail -
func (m *DB) GetUserByEmail(email string) (*models.User, UID, error) {
	return m.getUser("email", email)
}

// GetUserByID -
func (m *DB) GetUserByID(uid UID) (*models.User, UID, error) {
	return m.getUser("uid", uid)
}

// ----------------|

// UpdateUser -
func (m *DB) UpdateUser(nickname string, update *models.UserUpdate) error {
	tx := m.db.MustBegin()
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
		update.Fullname, update.Email, update.About, nickname)
	if err != nil {
		return err
	}
	return tx.Commit()
}
