package sqlite

import (
	"github.com/giicoo/GiicooAuth/internal/models"
)

func (sq *Sqlite) GetUserById(id int) (models.User, error) {
	stmt := "SELECT * FROM users WHERE user_id = ?;"

	user := models.User{}

	row := sq.db.QueryRow(stmt, id)
	err := row.Scan(&user.UserId, &user.Email, &user.HashPassword)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (sq *Sqlite) GetUserByEmail(email string) (models.User, error) {
	stmt := "SELECT * FROM users WHERE email = ?;"

	user := models.User{}

	row := sq.db.QueryRow(stmt, email)
	err := row.Scan(&user.UserId, &user.Email, &user.HashPassword)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (sq *Sqlite) CreateUser(email string, hash_password string) error {
	stmt := "INSERT INTO users (email, hash_password) VALUES (?, ?);"
	s, err := sq.db.Exec(stmt, email, hash_password)
	sq.log.Info(s)
	if err != nil {
		return err
	}
	return nil
}

//TODO: mutex
