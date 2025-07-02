package service

import (
	"time"

	"Ainotes/dto"
	"Ainotes/model"
	"Ainotes/repository"

	"github.com/google/uuid"
)

type NoteServiceImpl struct {
	noteRepo repository.NoteRepository
}

func NewNoteService(noteRepo repository.NoteRepository) NoteService {
	return &NoteServiceImpl{noteRepo: noteRepo}
}

func (svc *NoteServiceImpl) Create(req dto.NoteCreateRequest) (dto.NoteResponse, error) {
	note := model.Note{
		ID:         uuid.New(),
		ResourceID: req.ResourceID,
		UserID:     req.UserID,
		Title:      req.Title,
		Summary:    req.Summary,
		FullText:   req.FullText,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	if err := svc.noteRepo.Create(&note); err != nil {
		return dto.NoteResponse{}, err
	}
	return dto.NoteResponse{
		ID:         note.ID,
		ResourceID: note.ResourceID,
		UserID:     note.UserID,
		Title:      note.Title,
		Summary:    note.Summary,
		FullText:   note.FullText,
		CreatedAt:  note.CreatedAt,
		UpdatedAt:  note.UpdatedAt,
	}, nil
}

func (svc *NoteServiceImpl) FindByID(id uuid.UUID) (dto.NoteResponse, error) {
	note, err := svc.noteRepo.FindByID(id)
	if err != nil {
		return dto.NoteResponse{}, err
	}
	return dto.NoteResponse{
		ID:         note.ID,
		ResourceID: note.ResourceID,
		UserID:     note.UserID,
		Title:      note.Title,
		Summary:    note.Summary,
		FullText:   note.FullText,
		CreatedAt:  note.CreatedAt,
		UpdatedAt:  note.UpdatedAt,
	}, nil
}

func (svc *NoteServiceImpl) FindAll() ([]dto.NoteResponse, error) {
	notes, err := svc.noteRepo.FindAll()
	if err != nil {
		return nil, err
	}
	responses := make([]dto.NoteResponse, len(notes))
	for i, note := range notes {
		responses[i] = dto.NoteResponse{
			ID:         note.ID,
			ResourceID: note.ResourceID,
			UserID:     note.UserID,
			Title:      note.Title,
			Summary:    note.Summary,
			FullText:   note.FullText,
			CreatedAt:  note.CreatedAt,
			UpdatedAt:  note.UpdatedAt,
		}
	}
	return responses, nil
}

func (svc *NoteServiceImpl) Update(id uuid.UUID, req dto.NoteUpdateRequest) (dto.NoteResponse, error) {
    note, err := svc.noteRepo.FindByID(id)
    if err != nil {
        return dto.NoteResponse{}, err
    }
    note.Title = req.Title
    note.Summary = req.Summary
    note.FullText = req.FullText
    note.UpdatedAt = time.Now()
    if err := svc.noteRepo.Update(note); err != nil {
        return dto.NoteResponse{}, err
    }
    return dto.NoteResponse{
        ID:         note.ID,
        ResourceID: note.ResourceID,
        UserID:     note.UserID,
        Title:      note.Title,
        Summary:    note.Summary,
        FullText:   note.FullText,
        CreatedAt:  note.CreatedAt,
        UpdatedAt:  note.UpdatedAt,
    }, nil
}

func (svc *NoteServiceImpl) Delete(id uuid.UUID) error {
    return svc.noteRepo.Delete(id)
}