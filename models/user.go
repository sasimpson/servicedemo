package models

import "time"

//User is a user in the system
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Birthday  time.Time `json:"birthday"`
}

//UserModel is the interface for the User
type UserModel interface {
	All() (*[]User, error)
	Get(id int) (*User, error)
}
