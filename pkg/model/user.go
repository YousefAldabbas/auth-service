package model

type User struct {
	ID       int    `json:"id"`
	UUID     string    `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
