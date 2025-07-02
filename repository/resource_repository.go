package repository

import (
	"Ainotes/model"

	"github.com/google/uuid"
)

type ResourceRepository interface {
	Create(resource *model.Resource) error
	FindByID(id uuid.UUID) (*model.Resource, error)
	FindAll() ([]model.Resource, error)
	Update(resource *model.Resource) error
	Delete(id uuid.UUID) error
}
