package repository

import (
	"Ainotes/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuizRepositoryImpl struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) QuizRepository {
	return &QuizRepositoryImpl{db: db}
}

func (repo *QuizRepositoryImpl) Create(quiz *model.Quiz) error {
	return repo.db.Create(quiz).Error
}

func (repo *QuizRepositoryImpl) FindByID(id uuid.UUID) (*model.Quiz, error) {
	var quiz model.Quiz
	if err := repo.db.First(&quiz, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (repo *QuizRepositoryImpl) FindAll() ([]model.Quiz, error) {
	var quizzes []model.Quiz
	if err := repo.db.Find(&quizzes).Error; err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (repo *QuizRepositoryImpl) Update(quiz *model.Quiz) error {
	return repo.db.Save(quiz).Error
}

func (repo *QuizRepositoryImpl) Delete(id uuid.UUID) error {
	return repo.db.Delete(&model.Quiz{}, "id = ?", id).Error
}
