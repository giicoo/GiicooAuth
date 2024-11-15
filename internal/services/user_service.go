package services

import (
	"fmt"

	"github.com/giicoo/GiicooAuth/internal/models"
	errTools "github.com/giicoo/GiicooAuth/pkg/err_tools"
)

func (s *Services) CreateUser(email, password string) (models.UserResponse, error) {

	userYet, _ := s.GetUserByEmail(email)
	if userYet.Email == email {
		return models.UserResponse{}, errTools.WrapError(fmt.Errorf("user email %q already used", email), errTools.ErrEmailUsed)
	}

	hashPassword, err := s.hashTools.HashPassword(password)
	if err != nil {
		return models.UserResponse{}, err
	}
	user := models.User{
		Email:        email,
		HashPassword: hashPassword,
	}
	err = s.repo.CreateUser(user.Email, user.HashPassword)
	if err != nil {
		return models.UserResponse{}, err
	}

	userDB, err := s.repo.GetUserByEmail(email)
	userResponse := models.UserResponse{
		UserId: userDB.UserId,
		Email:  userDB.Email,
	}
	return userResponse, nil
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
