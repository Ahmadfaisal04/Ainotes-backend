package service

import (
	"time"

	"Ainotes/dto"
	"Ainotes/model"
	"Ainotes/repository"

	"github.com/google/uuid"
)

type ResourceServiceImpl struct {
	resourceRepo repository.ResourceRepository
}

func NewResourceService(resourceRepo repository.ResourceRepository) ResourceService {
	return &ResourceServiceImpl{resourceRepo: resourceRepo}
}

func (svc *ResourceServiceImpl) Create(req dto.ResourceCreateRequest) (dto.ResourceResponse, error) {
	resource := model.Resource{
		ID:            uuid.New(),
		UserID:        req.UserID,
		Type:          req.Type,
		SourceURL:     req.SourceURL,
		OriginalTitle: req.OriginalTitle,
		Status:        req.Status,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	if err := svc.resourceRepo.Create(&resource); err != nil {
		return dto.ResourceResponse{}, err
	}
	return dto.ResourceResponse{
		ID:            resource.ID,
		UserID:        resource.UserID,
		Type:          resource.Type,
		SourceURL:     resource.SourceURL,
		OriginalTitle: resource.OriginalTitle,
		Status:        resource.Status,
		CreatedAt:     resource.CreatedAt,
		UpdatedAt:     resource.UpdatedAt,
	}, nil
}

func (svc *ResourceServiceImpl) FindByID(id uuid.UUID) (dto.ResourceResponse, error) {
	resource, err := svc.resourceRepo.FindByID(id)
	if err != nil {
		return dto.ResourceResponse{}, err
	}
	return dto.ResourceResponse{
		ID:            resource.ID,
		UserID:        resource.UserID,
		Type:          resource.Type,
		SourceURL:     resource.SourceURL,
		OriginalTitle: resource.OriginalTitle,
		Status:        resource.Status,
		CreatedAt:     resource.CreatedAt,
		UpdatedAt:     resource.UpdatedAt,
	}, nil
}

func (svc *ResourceServiceImpl) FindAll() ([]dto.ResourceResponse, error) {
	resources, err := svc.resourceRepo.FindAll()
	if err != nil {
		return nil, err
	}
	responses := make([]dto.ResourceResponse, len(resources))
	for i, resource := range resources {
		responses[i] = dto.ResourceResponse{
			ID:            resource.ID,
			UserID:        resource.UserID,
			Type:          resource.Type,
			SourceURL:     resource.SourceURL,
			OriginalTitle: resource.OriginalTitle,
			Status:        resource.Status,
			CreatedAt:     resource.CreatedAt,
			UpdatedAt:     resource.UpdatedAt,
		}
	}
	return responses, nil
}

func (svc *ResourceServiceImpl) Update(id uuid.UUID, req dto.ResourceUpdateRequest) (dto.ResourceResponse, error) {
	resource, err := svc.resourceRepo.FindByID(id)
	if err != nil {
		return dto.ResourceResponse{}, err
	}
	resource.Type = req.Type
	resource.SourceURL = req.SourceURL
	resource.OriginalTitle = req.OriginalTitle
	resource.Status = req.Status
	resource.UpdatedAt = time.Now()
	if err := svc.resourceRepo.Update(resource); err != nil {
		return dto.ResourceResponse{}, err
	}
	return dto.ResourceResponse{
		ID:            resource.ID,
		UserID:        resource.UserID,
		Type:          resource.Type,
		SourceURL:     resource.SourceURL,
		OriginalTitle: resource.OriginalTitle,
		Status:        resource.Status,
		CreatedAt:     resource.CreatedAt,
		UpdatedAt:     resource.UpdatedAt,
	}, nil
}

func (svc *ResourceServiceImpl) Delete(id uuid.UUID) error {
	return svc.resourceRepo.Delete(id)
}
