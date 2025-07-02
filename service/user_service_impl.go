package service

import (
	"time"

	"Ainotes/dto"
	"Ainotes/model"
	"Ainotes/repository"

	"github.com/google/uuid"
)

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

func (svc *UserServiceImpl) Create(req dto.UserCreateRequest) (dto.UserResponse, error) {
	user := model.User{
		ID:        uuid.New(),
		GoogleID:  req.GoogleID,
		Email:     req.Email,
		FullName:  req.FullName,
		AvatarURL: req.AvatarURL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := svc.userRepo.Create(&user); err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{
		ID:        user.ID,
		GoogleID:  user.GoogleID,
		Email:     user.Email,
		FullName:  user.FullName,
		AvatarURL: user.AvatarURL,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (svc *UserServiceImpl) FindByID(id uuid.UUID) (dto.UserResponse, error) {
	user, err := svc.userRepo.FindByID(id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{
		ID:        user.ID,
		GoogleID:  user.GoogleID,
		Email:     user.Email,
		FullName:  user.FullName,
		AvatarURL: user.AvatarURL,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (svc *UserServiceImpl) FindAll() ([]dto.UserResponse, error) {
	users, err := svc.userRepo.FindAll()
	if err != nil {
		return nil, err
	}
	responses := make([]dto.UserResponse, len(users))
	for i, user := range users {
		responses[i] = dto.UserResponse{
			ID:        user.ID,
			GoogleID:  user.GoogleID,
			Email:     user.Email,
			FullName:  user.FullName,
			AvatarURL: user.AvatarURL,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}
	return responses, nil
}

func (svc *UserServiceImpl) Update(id uuid.UUID, req dto.UserUpdateRequest) (dto.UserResponse, error) {
	user, err := svc.userRepo.FindByID(id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	user.FullName = req.FullName
	user.AvatarURL = req.AvatarURL
	user.UpdatedAt = time.Now()
	if err := svc.userRepo.Update(user); err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{
		ID:        user.ID,
		GoogleID:  user.GoogleID,
		Email:     user.Email,
		FullName:  user.FullName,
		AvatarURL: user.AvatarURL,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (svc *UserServiceImpl) Delete(id uuid.UUID) error {
	return svc.userRepo.Delete(id)
}