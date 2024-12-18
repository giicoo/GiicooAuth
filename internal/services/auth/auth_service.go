package auth

import (
	"github.com/giicoo/GiicooAuth/internal/config"
	"github.com/giicoo/GiicooAuth/internal/repository"
	"github.com/giicoo/GiicooAuth/pkg/errTools"
	"github.com/giicoo/GiicooAuth/pkg/hash"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	cfg  *config.Config
	log  *logrus.Logger
	repo repository.Repo

	hashTools  hash.HashTools
	jwtManager *JwtManager
}

func NewAuthService(cfg *config.Config, log *logrus.Logger, repo repository.Repo) *AuthService {
	return &AuthService{
		cfg:  cfg,
		log:  log,
		repo: repo,

		hashTools:  hash.NewHashTools(),
		jwtManager: NewJwtManger(cfg.JWT.PathToKey),
	}
}

func (as *AuthService) Login(email string, password string) (string, string, error) {

	userDB, err := as.repo.GetUserByEmail(email)
	if err != nil {
		return "", "", errTools.WrapError(err, errTools.ErrEmailNotReg)
	}

	if !as.hashTools.CheckPasswordHash(password, userDB.HashPassword) {
		return "", "", errTools.WrapError(err, errTools.ErrWrongPassword)
	}

	access, refresh, err := as.GenerateTokens(userDB.UserId)
	if err != nil {
		return "", "", err
	}

	if err := as.repo.SaveRefreshTokenToDB(userDB.UserId, refresh); err != nil {
		return "", "", err
	}
	return access, refresh, nil

}

func (as *AuthService) GenerateNewAccessToken(userID int) (string, error) {
	return as.jwtManager.GenerateAccessToken(userID, as.cfg.JWT.Access.Time)
}

func (as *AuthService) GenerateTokens(userID int) (string, string, error) {
	access, refresh, err := as.jwtManager.GenerateTokens(userID, as.cfg.JWT.Access.Time, as.cfg.JWT.Refresh.Time)
	if err != nil {
		return "", "", err
	}

	if err := as.repo.SaveRefreshTokenToDB(userID, refresh); err != nil {
		return "", "", err
	}
	return access, refresh, nil
}

func (as *AuthService) ValidateAccessToken(accessToken string) (int, error) {
	userID, err := as.jwtManager.ValidateAccessToken(accessToken)
	if err != nil {
		return 0, errTools.WrapError(err, errTools.ErrInvalidAccessToken)
	}
	return userID, nil
}

func (as *AuthService) ValidateRefreshToken(refreshToken string) (int, error) {
	userID, err := as.jwtManager.ValidateRefreshToken(refreshToken)
	if err != nil {
		return 0, errTools.WrapError(err, errTools.ErrInvalidRefreshToken)
	}
	return userID, nil
}
