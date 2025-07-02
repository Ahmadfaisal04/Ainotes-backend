package service

import (
	"Ainotes/dto"

	"github.com/google/uuid"
)

type ResourceService interface {
	Create(req dto.ResourceCreateRequest) (dto.ResourceResponse, error)
	FindByID(id uuid.UUID) (dto.ResourceResponse, error)
	FindAll() ([]dto.ResourceResponse, error)
	Update(id uuid.UUID, req dto.ResourceUpdateRequest) (dto.ResourceResponse, error)
	Delete(id uuid.UUID) error
}