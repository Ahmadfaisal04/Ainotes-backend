package dto

import (
	"time"

	"github.com/google/uuid"
)

type ResourceResponse struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Type          string    `json:"type"`
	SourceURL     string    `json:"source_url"`
	OriginalTitle string    `json:"original_title"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}