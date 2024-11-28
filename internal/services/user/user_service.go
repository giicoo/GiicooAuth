package user

import (
	"fmt"

	"github.com/giicoo/GiicooAuth/internal/config"
	"github.com/giicoo/GiicooAuth/internal/models"
	"github.com/giicoo/GiicooAuth/internal/repository"
	"github.com/giicoo/GiicooAuth/pkg/errTools"
	"github.com/giicoo/GiicooAuth/pkg/hash"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	cfg  *config.Config
	log  *logrus.Logger
	repo repository.Repo

	hashTools hash.HashTools
}

func NewUserService(cfg *config.Config, log *logrus.Logger, repo repository.Repo) *UserService {
	return &UserService{
		cfg:  cfg,
		log:  log,
		repo: repo,

		hashTools: hash.NewHashTools(),
	}
}

func (us *UserService) CreateUser(email, password string) (models.UserResponse, error) {

	userYet, _ := us.repo.GetUserByEmail(email)
	if userYet.Email == email {
		return models.UserResponse{}, errTools.WrapError(fmt.Errorf("user email %q already used", email), errTools.ErrEmailUsed)
	}

	hashPassword, err := us.hashTools.HashPassword(password)
	if err != nil {
		return models.UserResponse{}, err
	}
	user := models.User{
		Email:        email,
		HashPassword: hashPassword,
	}
	err = us.repo.CreateUser(user.Email, user.HashPassword)
	if err != nil {
		return models.UserResponse{}, err
	}

	userDB, err := us.repo.GetUserByEmail(email)
	if err != nil {
		return models.UserResponse{}, err
	}
	userResponse := models.UserResponse{
		UserId: userDB.UserId,
		Email:  userDB.Email,
	}
	return userResponse, nil
}

func (us *UserService) GetUserByEmail(email string) (models.User, error) {
	user, err := us.repo.GetUserByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (us *UserService) GetUserById(id int) (models.User, error) {
	user, err := us.repo.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
