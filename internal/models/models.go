package models

type User struct {
	UserId       int    `json:"user_id"`
	Email        string `json:"email"`
	HashPassword string `json:"hash_password"`
}

type UserRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
}

type JwtRequest struct {
	JwtToken string `json:"jwt" validate:"required"`
}

type JwtResponse struct {
	JwtToken string `json:"jwt"`
}

type ErrorResponse struct {
	Err string `json:"error"`
}
