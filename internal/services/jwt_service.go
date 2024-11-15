package services

import (
	"strings"
	"time"

	"github.com/giicoo/GiicooAuth/internal/models"
	errTools "github.com/giicoo/GiicooAuth/pkg/err_tools"
)

func (s *Services) GenerateJWT(email, password string) (models.JwtResponse, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return models.JwtResponse{}, err
	}

	if s.hashTools.CheckPasswordHash(password, user.HashPassword) {
		jwtToken, err := s.jwtTools.NewJWT(user.UserId, user.Email, 60*time.Minute)
		if err != nil {
			return models.JwtResponse{}, err
		}
		jwtResponse := models.JwtResponse{JwtToken: jwtToken}
		return jwtResponse, nil
	} else {
		return models.JwtResponse{}, errTools.WrapError(err, errTools.ErrWrongPassword)
	}
}

func (s *Services) CheckJWT(jwtToken string) (models.UserResponse, error) {
	uid, email, err := s.jwtTools.ParseJWT(jwtToken)
	if err != nil {
		switch {
		case strings.HasPrefix(err.Error(), "token is expired"):
			{
				return models.UserResponse{}, errTools.WrapError(err, errTools.ErrInvalidTokenExpired)
			}
		default:
			{
				return models.UserResponse{}, errTools.WrapError(err, errTools.ErrInvalidTokenSyntax)
			}
		}

	}

	userResponse := models.UserResponse{
		UserId: uid,
		Email:  email,
	}

	return userResponse, nil
}
