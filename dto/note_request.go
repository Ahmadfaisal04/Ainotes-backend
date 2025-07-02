package dto

import "github.com/google/uuid"

type NoteCreateRequest struct {
	ResourceID uuid.UUID `json:"resource_id" binding:"required"`
	UserID     uuid.UUID `json:"user_id" binding:"required"`
	Title      string    `json:"title" binding:"required"`
	Summary    string    `json:"summary"`
	FullText   string    `json:"full_text"`
}

type NoteUpdateRequest struct {
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	FullText string `json:"full_text"`
}