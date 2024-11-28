package models

import "github.com/golang-jwt/jwt"

type User struct {
	UserId       int    `json:"user_id"`
	Email        string `json:"email"`
	HashPassword string `json:"hash_password"`
	RefreshToken string `json:"refresh_token"`
}

type UserRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
}

type AccessTokenRequest struct {
	AccessToken string `json:"access_token" validate:"required"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type ValidAccessTokenResponse struct {
	Valid  bool `json:"valid"`
	UserID int  `json:"user_id"`
}

type InvalidAccessTokenResponse struct {
	Valid       bool   `json:"valid"`
	AccessToken string `json:"new_access_token"`
}

type ErrorResponse struct {
	Err string `json:"error"`
}

type Claims struct {
	jwt.StandardClaims
	UserId    int    `json:"id"`
	TokenType string `json:"token_type"`
}
