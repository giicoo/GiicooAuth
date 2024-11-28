package repository

import "github.com/giicoo/GiicooAuth/internal/models"

type Repo interface {
	InitDB() error
	GetUserById(id int) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	CreateUser(email string, hash_password string) error

	// SaveRefreshTokenToDB(userID int, refreshToken string)
}
