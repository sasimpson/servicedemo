package mock

import "github.com/sasimpson/servicedemo/models"

type UserMock struct {
	Users *[]models.User
	Error error
}

func (m *UserMock) All() (*[]models.User, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Users, nil
}

func (m *UserMock) Get(id int) (*models.User, error) {
	panic("not implemented")
}
