package service

import (
	"Ainotes/dto"

	"github.com/google/uuid"
)

type NoteService interface {
	Create(req dto.NoteCreateRequest) (dto.NoteResponse, error)
	FindByID(id uuid.UUID) (dto.NoteResponse, error)
	FindAll() ([]dto.NoteResponse, error)
	Update(id uuid.UUID, req dto.NoteUpdateRequest) (dto.NoteResponse, error)
	Delete(id uuid.UUID) error
}