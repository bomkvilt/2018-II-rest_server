package database

import (
	"Forum/app/generated/models"
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
