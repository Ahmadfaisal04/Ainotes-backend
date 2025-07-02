package service

import (
	"Ainotes/dto"

	"github.com/google/uuid"
)

type UserService interface {
	Create(req dto.UserCreateRequest) (dto.UserResponse, error)
	FindByID(id uuid.UUID) (dto.UserResponse, error)
	FindAll() ([]dto.UserResponse, error)
	Update(id uuid.UUID, req dto.UserUpdateRequest) (dto.UserResponse, error)
	Delete(id uuid.UUID) error
}