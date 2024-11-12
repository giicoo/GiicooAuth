package services

import (
	"errors"

	"github.com/giicoo/GiicooAuth/internal/models"
)

func (s *Services) CreateUser(email, password string) error {

	userYet, _ := s.GetUserByEmail(email)
	if userYet.Email == email {
		return errors.New("User already be")
	}

	hashPassword, err := s.hashTools.HashPassword(password)
	if err != nil {
		return err
	}
	user := models.User{
		Email:        email,
		HashPassword: hashPassword,
	}
	return s.repo.CreateUser(user.Email, user.HashPassword)
}

func (s *Services) GetUserByEmail(email string) (models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *Services) GetUserById(id int) (models.User, error) {
	user, err := s.repo.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
