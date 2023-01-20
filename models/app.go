package models

import "time"

//App is a registered application that can use the API
type App struct {
	ID         int
	Name       string
	Key        string
	Expiration int32
	CreatedAt  time.Time
	User       *User
}

//AppModel is the interface definition
type AppModel interface {
	Get(int) (*App, error)
	GetByKey(string) (*App, error)
	New(*App) (int, error)
	Delete(int) error
}
