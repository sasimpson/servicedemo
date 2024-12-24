package mock

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sasimpson/servicedemo/models"
)

// Product
// Implements ProductDataInterface
type Product struct {
	Products *[]models.Product
	Product  *models.Product
	Error    error
}

func (m Product) All() (*[]models.Product, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Products, nil
}

func (m Product) Get(_ uuid.UUID) (*models.Product, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	if m.Products == nil {
		return nil, fmt.Errorf("product %w", models.ErrNotFound)
	}
	return m.Product, nil
}

func (m Product) New(_ *models.Product) (*models.Product, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	return m.Product, nil
}

func (m Product) Update(_ *models.Product) (*models.Product, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	if m.Products == nil {
		return nil, fmt.Errorf("product %v, need product to update", models.ErrNilValue)
	}
	return m.Product, nil
}

func (m Product) Delete(_ uuid.UUID) error {
	if m.Error != nil {
		return m.Error
	}
	return nil
}
