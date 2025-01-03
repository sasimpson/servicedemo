package models

import (
	"errors"
	"time"
)

// Env structure holds our environment settings
type Env struct {
}

var (
	//ErrNotFound is the generic "not found" error for being wrapped
	ErrNotFound = errors.New("not found")
	//ErrAlreadyExists is the generic "already exists" error for being wrapped
	ErrAlreadyExists = errors.New("exists already")
)

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
