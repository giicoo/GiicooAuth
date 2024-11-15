package sqlite

import (
	"os"

	"github.com/giicoo/GiicooAuth/internal/models"
)

func (sq *Sqlite) GetUserById(id int) (models.User, error) {

	file, err := os.ReadFile(sq.cfg.DB.PathToSQL + "get_user_by_id.sql")
	if err != nil {
		return models.User{}, err
	}
	user := models.User{}

	stmt := string(file)
	row := sq.db.QueryRow(stmt, id)
	err = row.Scan(&user.UserId, &user.Email, &user.HashPassword)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (sq *Sqlite) GetUserByEmail(email string) (models.User, error) {

	file, err := os.ReadFile(sq.cfg.DB.PathToSQL + "get_user_by_email.sql")
	if err != nil {
		return models.User{}, err
	}
	user := models.User{}

	stmt := string(file)
	row := sq.db.QueryRow(stmt, email)
	err = row.Scan(&user.UserId, &user.Email, &user.HashPassword)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (sq *Sqlite) CreateUser(email string, hash_password string) error {

	file, err := os.ReadFile(sq.cfg.DB.PathToSQL + "insert_user.sql")
	if err != nil {
		return err
	}
	stmt := string(file)
	s, err := sq.db.Exec(stmt, email, hash_password)
	sq.log.Info(s)
	if err != nil {
		return err
	}
	return nil
}

//TODO: mutex
