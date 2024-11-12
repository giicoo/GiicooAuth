package models

type User struct {
	UserId       int    `json:"user_id"`
	Email        string `json:"email"`
	HashPassword string `json:"hash_password"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserResponse struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
}
type JwtResponse struct {
	JwtToken string `json:"jwt"`
}
