package jwt_auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	Login     string `json:"login"`
	UserId    int    `json:"id"`
	TypeToken string `json:"type_token"`
}
type JWT struct {
	key string
}

func NewJWT(key string) *JWT {
	return &JWT{
		key: key,
	}
}
func (j *JWT) NewJWT(id int, login string, d time.Duration) (string, error) {
	// generate jwt token by user login with
	// + ExpiresAt - time for access
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(d).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
		},
		Login:  login,
		UserId: id,
	})
	return token.SignedString([]byte(j.key))
}

func (j *JWT) ParseJWT(tk string) (int, string, error) {
	// parse token
	token, err := jwt.ParseWithClaims(tk, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// func to check method how hash token and send key for this token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.key), nil
	})
	if err != nil {
		return 0, "", err
	}

	// check valid token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserId, claims.Login, nil
	}

	return 0, "", errors.New("Invalid token")
}
