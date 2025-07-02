package repository

import (
	"Ainotes/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) Create(user *model.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepositoryImpl) FindByID(id uuid.UUID) (*model.User, error) {
	var user model.User
	if err := repo.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepositoryImpl) FindAll() ([]model.User, error) {
	var users []model.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepositoryImpl) Update(user *model.User) error {
	return repo.db.Save(user).Error
}

func (repo *UserRepositoryImpl) Delete(id uuid.UUID) error {
	return repo.db.Delete(&model.User{}, "id = ?", id).Error
}
