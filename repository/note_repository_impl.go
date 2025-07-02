package repository

import (
	"Ainotes/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NoteRepositoryImpl struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &NoteRepositoryImpl{db: db}
}

func (repo *NoteRepositoryImpl) Create(note *model.Note) error {
	return repo.db.Create(note).Error
}

func (repo *NoteRepositoryImpl) FindByID(id uuid.UUID) (*model.Note, error) {
	var note model.Note
	if err := repo.db.First(&note, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &note, nil
}

func (repo *NoteRepositoryImpl) FindAll() ([]model.Note, error) {
	var notes []model.Note
	if err := repo.db.Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func (repo *NoteRepositoryImpl) Update(note *model.Note) error {
	return repo.db.Save(note).Error
}

func (repo *NoteRepositoryImpl) Delete(id uuid.UUID) error {
	return repo.db.Delete(&model.Note{}, "id = ?", id).Error
}
