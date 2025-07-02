package service

import (
	"Ainotes/dto"
	"Ainotes/model"
	"Ainotes/repository"

	"github.com/google/uuid"
)

type FlashcardServiceImpl struct {
	flashcardRepo repository.FlashcardRepository
}

func NewFlashcardService(flashcardRepo repository.FlashcardRepository) FlashcardService {
	return &FlashcardServiceImpl{flashcardRepo: flashcardRepo}
}

func (svc *FlashcardServiceImpl) Create(req dto.FlashcardCreateRequest) (dto.FlashcardResponse, error) {
	flashcard := model.Flashcard{
		ID:        uuid.New(),
		NoteID:    req.NoteID,
		FrontText: req.FrontText,
		BackText:  req.BackText,
	}
	if err := svc.flashcardRepo.Create(&flashcard); err != nil {
		return dto.FlashcardResponse{}, err
	}
	return dto.FlashcardResponse{
		ID:        flashcard.ID,
		NoteID:    flashcard.NoteID,
		FrontText: flashcard.FrontText,
		BackText:  flashcard.BackText,
	}, nil
}

func (svc *FlashcardServiceImpl) FindByID(id uuid.UUID) (dto.FlashcardResponse, error) {
	flashcard, err := svc.flashcardRepo.FindByID(id)
	if err != nil {
		return dto.FlashcardResponse{}, err
	}
	return dto.FlashcardResponse{
		ID:        flashcard.ID,
		NoteID:    flashcard.NoteID,
		FrontText: flashcard.FrontText,
		BackText:  flashcard.BackText,
	}, nil
}

func (svc *FlashcardServiceImpl) FindAll() ([]dto.FlashcardResponse, error) {
	flashcards, err := svc.flashcardRepo.FindAll()
	if err != nil {
		return nil, err
	}
	responses := make([]dto.FlashcardResponse, len(flashcards))
	for i, flashcard := range flashcards {
		responses[i] = dto.FlashcardResponse{
			ID:        flashcard.ID,
			NoteID:    flashcard.NoteID,
			FrontText: flashcard.FrontText,
			BackText:  flashcard.BackText,
		}
	}
	return responses, nil
}

func (svc *FlashcardServiceImpl) Update(id uuid.UUID, req dto.FlashcardUpdateRequest) (dto.FlashcardResponse, error) {
	flashcard, err := svc.flashcardRepo.FindByID(id)
	if err != nil {
		return dto.FlashcardResponse{}, err
	}
	flashcard.FrontText = req.FrontText
	flashcard.BackText = req.BackText
	if err := svc.flashcardRepo.Update(flashcard); err != nil {
		return dto.FlashcardResponse{}, err
	}
	return dto.FlashcardResponse{
		ID:        flashcard.ID,
		NoteID:    flashcard.NoteID,
		FrontText: flashcard.FrontText,
		BackText:  flashcard.BackText,
	}, nil
}

func (svc *FlashcardServiceImpl) Delete(id uuid.UUID) error {
	return svc.flashcardRepo.Delete(id)
}