package product

import "time"

// Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Models slice of Model
type Models []*Model

type Storage interface {
	Create(*Model) error
	Updated(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error
}
