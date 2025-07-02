package service

import (
	"Ainotes/dto"

	"github.com/google/uuid"
)

type FlashcardService interface {
	Create(req dto.FlashcardCreateRequest) (dto.FlashcardResponse, error)
	FindByID(id uuid.UUID) (dto.FlashcardResponse, error)
	FindAll() ([]dto.FlashcardResponse, error)
	Update(id uuid.UUID, req dto.FlashcardUpdateRequest) (dto.FlashcardResponse, error)
	Delete(id uuid.UUID) error
}