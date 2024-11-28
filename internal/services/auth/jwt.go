package auth

import (
	"errors"
	"os"
	"time"

	"github.com/giicoo/GiicooAuth/internal/models"
	"github.com/golang-jwt/jwt"
)

type JwtManager struct {
	pathToKey string
}

func NewJwtManger(pathToKey string) *JwtManager {
	return &JwtManager{
		pathToKey: pathToKey,
	}
}
func (jm *JwtManager) GenerateTokens(userID int, accessTokenTime int, refreshTokenTime int) (string, string, error) {
	accessToken, err := jm.GenerateAccessToken(userID, accessTokenTime)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jm.GenerateRefreshToken(userID, refreshTokenTime)
	if err != nil {
		return "", "", err
	}
	// as.repo.SaveRefreshTokenToDB(userID, refreshToken)

	return accessToken, refreshToken, nil
}

func (jm *JwtManager) GenerateAccessToken(userID int, accessTokenTime int) (string, error) {
	tokenType := "access"

	claims := models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(accessTokenTime)).Unix(),
			Issuer:    "auth.service",
		},
		UserId:    userID,
		TokenType: tokenType,
	}
	key, err := os.ReadFile(jm.pathToKey)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}

func (jm *JwtManager) GenerateRefreshToken(userID int, refreshTokenTime int) (string, error) {
	tokenType := "refresh"

	claims := models.Claims{
		UserId:    userID,
		TokenType: tokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(refreshTokenTime)).Unix(),
			Issuer:    "auth.service",
		},
	}
	key, err := os.ReadFile(jm.pathToKey)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}

func (jm *JwtManager) ValidateAccessToken(tokenStr string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method in auth token")
		}
		key, err := os.ReadFile(jm.pathToKey)

		return key, err
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*models.Claims)

	if !ok || !token.Valid || claims.UserId == 0 || claims.TokenType != "access" {
		return 0, errors.New("invalid token: authentication failed")
	}
	return claims.UserId, nil
	//TODO: return all claims
}

func (jm *JwtManager) ValidateRefreshToken(tokenStr string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, errors.New("Unexpected signing method in auth token")
		}
		key, err := os.ReadFile(jm.pathToKey)

		return key, err
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*models.Claims)

	if !ok || !token.Valid || claims.UserId == 0 || claims.TokenType != "refresh" {
		return 0, errors.New("invalid token: authentication failed")
	}
	return claims.UserId, nil
}
