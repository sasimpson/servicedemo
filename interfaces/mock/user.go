package mock

import "github.com/sasimpson/servicedemo/models"

type UserMock struct{}

func (m *UserMock) All() (*[]models.User, error) {
	panic("not implemented")
}

func (m *UserMock) Get(id int) (*models.User, error) {
	panic("not implemented")
}
