package mock

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sasimpson/servicedemo/models"
)

// User structure contains all the things that each of the mocks might hand back.
// Implements UserDataInterface
type User struct {
	Users *[]models.User
	User  *models.User
	Error error
}

// All mocks All function in our interface
func (m *User) All() (*[]models.User, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Users, nil
}

// Get mocks Get function in our interface
func (m *User) Get(_ uuid.UUID) (*models.User, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	if m.User == nil {
		return nil, fmt.Errorf("user %w", models.ErrNotFound)
	}
	return m.User, nil
}

// New mocks New function in our interface
func (m *User) New(*models.User) (*models.User, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	return m.User, nil
}
