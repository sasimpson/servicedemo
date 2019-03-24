package models

import "time"

type Env struct {
}
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Birthday  time.Time `json:"birthday"`
}

type UserModel interface {
	All() (*[]User, error)
	Get(id int) (*User, error)
}
