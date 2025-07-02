package repository

import (
	"Ainotes/model"

	"github.com/google/uuid"
)

type QuizRepository interface {
	Create(quiz *model.Quiz) error
	FindByID(id uuid.UUID) (*model.Quiz, error)
	FindAll() ([]model.Quiz, error)
	Update(quiz *model.Quiz) error
	Delete(id uuid.UUID) error
}
