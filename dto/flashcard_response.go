package dto

import "github.com/google/uuid"

type FlashcardResponse struct {
	ID        uuid.UUID `json:"id"`
	NoteID    uuid.UUID `json:"note_id"`
	FrontText string    `json:"front_text"`
	BackText  string    `json:"back_text"`
}