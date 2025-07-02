package repository

import (
	"Ainotes/model"

	"github.com/google/uuid"
)

type FlashcardRepository interface {
	Create(flashcard *model.Flashcard) error
	FindByID(id uuid.UUID) (*model.Flashcard, error)
	FindAll() ([]model.Flashcard, error)
	Update(flashcard *model.Flashcard) error
	Delete(id uuid.UUID) error
}
