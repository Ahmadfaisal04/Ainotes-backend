package repository

import (
	"Ainotes/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FlashcardRepositoryImpl struct {
	db *gorm.DB
}

func NewFlashcardRepository(db *gorm.DB) FlashcardRepository {
	return &FlashcardRepositoryImpl{db: db}
}

func (repo *FlashcardRepositoryImpl) Create(flashcard *model.Flashcard) error {
	return repo.db.Create(flashcard).Error
}

func (repo *FlashcardRepositoryImpl) FindByID(id uuid.UUID) (*model.Flashcard, error) {
	var flashcard model.Flashcard
	if err := repo.db.First(&flashcard, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &flashcard, nil
}

func (repo *FlashcardRepositoryImpl) FindAll() ([]model.Flashcard, error) {
	var flashcards []model.Flashcard
	if err := repo.db.Find(&flashcards).Error; err != nil {
		return nil, err
	}
	return flashcards, nil
}

func (repo *FlashcardRepositoryImpl) Update(flashcard *model.Flashcard) error {
	return repo.db.Save(flashcard).Error
}

func (repo *FlashcardRepositoryImpl) Delete(id uuid.UUID) error {
	return repo.db.Delete(&model.Flashcard{}, "id = ?", id).Error
}