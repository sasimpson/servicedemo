package models

import (
	"github.com/google/uuid"
	"time"
)

// User is how the user is represented
type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Birthday  time.Time `json:"birthday"`
	Email     string    `json:"email"`

	BaseModel
}

// UserDataInterface is the interface definition for user datastore interfaces
type UserDataInterface interface {
	All() (*[]User, error)
	Get(id uuid.UUID) (*User, error)
	New(*User) (*User, error)
}
