package models

import "time"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Birthday  time.Time
}
