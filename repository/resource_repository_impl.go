package repository

import (
	"Ainotes/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResourceRepositoryImpl struct {
	db *gorm.DB
}

func NewResourceRepository(db *gorm.DB) ResourceRepository {
	return &ResourceRepositoryImpl{db: db}
}

func (repo *ResourceRepositoryImpl) Create(resource *model.Resource) error {
	return repo.db.Create(resource).Error
}

func (repo *ResourceRepositoryImpl) FindByID(id uuid.UUID) (*model.Resource, error) {
	var resource model.Resource
	if err := repo.db.First(&resource, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &resource, nil
}

func (repo *ResourceRepositoryImpl) FindAll() ([]model.Resource, error) {
	var resources []model.Resource
	if err := repo.db.Find(&resources).Error; err != nil {
		return nil, err
	}
	return resources, nil
}

func (repo *ResourceRepositoryImpl) Update(resource *model.Resource) error {
	return repo.db.Save(resource).Error
}

func (repo *ResourceRepositoryImpl) Delete(id uuid.UUID) error {
	return repo.db.Delete(&model.Resource{}, "id = ?", id).Error
}
