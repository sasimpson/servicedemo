package models

import (
	"time"
)

// User is how the user is represented
type User struct {
	ID        int       `json:"id" required:"true" description:"Unique ID"`
	FirstName string    `json:"first_name" description:"User first name"`
	LastName  string    `json:"last_name" description:"User Last name"`
	Birthday  time.Time `json:"birthday" description:"User birthday"`
	Email     string    `json:"email" required:"true" description:"User email address"`
	_         struct{}  `title:"User" description:"User Model"`
}

// UserModel is the interface definition for user datastore interfaces
type UserModel interface {
	All() (*[]User, error)
	Get(id int) (*User, error)
	New(*User) (*User, error)
}
