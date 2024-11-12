package services

import (
	"github.com/giicoo/GiicooAuth/internal/config"
	"github.com/giicoo/GiicooAuth/internal/repository"
	"github.com/giicoo/GiicooAuth/pkg/hash"
	"github.com/giicoo/GiicooAuth/pkg/jwt_auth"
	"github.com/sirupsen/logrus"
)

type Services struct {
	cfg  *config.Config
	log  *logrus.Logger
	repo repository.Repo

	hashTools *hash.Hash
	jwtTools  *jwt_auth.JWT
}

func NewServices(cfg *config.Config, log *logrus.Logger, repo repository.Repo) *Services {
	return &Services{
		cfg:       cfg,
		log:       log,
		repo:      repo,
		hashTools: hash.NewHashTools(),
		jwtTools:  jwt_auth.NewJWT("foo"),
	}
}
