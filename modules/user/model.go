package user

import "database/sql"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserProfileCheck struct {
	Id       int            `json:"id"`
	Username string         `json:"username"`
	Bio      sql.NullString `json:"bio"`
	Location sql.NullString `json:"location"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	UserCollection []User
)
