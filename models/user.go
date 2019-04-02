package models

import (
	"errors"
	"time"
)

//User is how the user is represented
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Birthday  time.Time `json:"birthday"`
	Email     string    `json:"email"`
}

//UserModel is the interface definition for user datastore interfaces
type UserModel interface {
	All() (*[]User, error)
	Get(id int) (*User, error)
	New(*User) (*User, error)
}

var (
	//ErrUserNotFound for when the database cannot find the user
	ErrUserNotFound = errors.New("User Not Found")
	//ErrUserExists for when the exists on a POST
	ErrUserExists = errors.New("User exists already")
)
