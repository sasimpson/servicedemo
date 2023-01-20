package mock

import (
	"fmt"
	"github.com/sasimpson/servicedemo/models"
)

// UserMock structure contains all of the things that each of the mocks might hand back.
type UserMock struct {
	Users *[]models.User
	User  *models.User
	Error error
}

// All mocks the All function in our interface
func (m *UserMock) All() (*[]models.User, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Users, nil
}

// Get mocks the Get function in our interface
func (m *UserMock) Get(_ int) (*models.User, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	if m.User == nil {
		return nil, fmt.Errorf("user %w", models.ErrNotFound)
	}
	return m.User, nil
}

// New mocks the New function in our interface
func (m *UserMock) New(*models.User) (*models.User, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	return m.User, nil
}
