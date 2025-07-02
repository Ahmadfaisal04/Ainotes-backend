package dto

import (
	"time"

	"github.com/google/uuid"
)

type NoteResponse struct {
	ID         uuid.UUID `json:"id"`
	ResourceID uuid.UUID `json:"resource_id"`
	UserID     uuid.UUID `json:"user_id"`
	Title      string    `json:"title"`
	Summary    string    `json:"summary"`
	FullText   string    `json:"full_text"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}