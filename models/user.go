package models

import "time"

type Env struct {
}
type User struct {
	ID        int
	FirstName string
	LastName  string
	Birthday  time.Time
}

type UserModel interface {
	All() (*[]User, error)
	Get(id int) (*User, error)
}
