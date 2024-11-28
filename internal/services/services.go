package services

import (
	"github.com/giicoo/GiicooAuth/internal/config"
	"github.com/giicoo/GiicooAuth/internal/repository"
	"github.com/giicoo/GiicooAuth/internal/services/auth"
	"github.com/giicoo/GiicooAuth/internal/services/user"
	"github.com/sirupsen/logrus"
)

type Services struct {
	AuthService *auth.AuthService
	UserService *user.UserService
}

func NewServices(cfg *config.Config, log *logrus.Logger, repo repository.Repo) *Services {
	return &Services{
		AuthService: auth.NewAuthService(cfg, log, repo),
		UserService: user.NewUserService(cfg, log, repo),
	}
}
