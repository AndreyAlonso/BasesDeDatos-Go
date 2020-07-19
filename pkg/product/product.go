package product

import "time"

// Models slice of Model
type Models []*Model

// Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Storage interface {
	Migrate() error
	Create(*Model) error
	GetAll() (Models, error)
}

// Serivce of product
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate product
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

// Create es usado para crear un producto
func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}

// GetAll es usado para obtener todos los productos
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}
