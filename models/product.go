package models

import (
	"github.com/google/uuid"
)

// Product is how a product is represented
type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Quantity    int       `json:"quantity"`

	BaseModel
}

// ProductDataInterface is the interface definition for the datastore interfaces
type ProductDataInterface interface {
	All() (*[]Product, error)
	Get(id uuid.UUID) (*Product, error)
	New(*Product) (*Product, error)
	Update(*Product) (*Product, error)
	Delete(id uuid.UUID) error
}
