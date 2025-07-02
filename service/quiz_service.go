package service

import (
	"Ainotes/dto"

	"github.com/google/uuid"
)

type QuizService interface {
	Create(req dto.QuizCreateRequest) (dto.QuizResponse, error)
	FindByID(id uuid.UUID) (dto.QuizResponse, error)
	FindAll() ([]dto.QuizResponse, error)
	Update(id uuid.UUID, req dto.QuizUpdateRequest) (dto.QuizResponse, error)
	Delete(id uuid.UUID) error
}