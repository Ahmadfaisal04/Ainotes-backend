package dto

import "github.com/google/uuid"

type FlashcardCreateRequest struct {
	NoteID    uuid.UUID `json:"note_id" binding:"required"`
	FrontText string    `json:"front_text" binding:"required"`
	BackText  string    `json:"back_text" binding:"required"`
}

type FlashcardUpdateRequest struct {
	FrontText string `json:"front_text"`
	BackText  string `json:"back_text"`
}