package repository

import (
	"Ainotes/model"

	"github.com/google/uuid"
)

type NoteRepository interface {
	Create(note *model.Note) error
	FindByID(id uuid.UUID) (*model.Note, error)
	FindAll() ([]model.Note, error)
	Update(note *model.Note) error
	Delete(id uuid.UUID) error
}
