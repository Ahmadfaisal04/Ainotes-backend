package repository

import (
	"Ainotes/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByID(id uuid.UUID) (*model.User, error)
	FindAll() ([]model.User, error)
	Update(user *model.User) error
	Delete(id uuid.UUID) error
}
