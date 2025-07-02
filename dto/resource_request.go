package dto

import "github.com/google/uuid"

type ResourceCreateRequest struct {
	UserID        uuid.UUID `json:"user_id" binding:"required"`
	Type          string    `json:"type" binding:"required"`
	SourceURL     string    `json:"source_url"`
	OriginalTitle string    `json:"original_title"`
	Status        string    `json:"status"`
}

type ResourceUpdateRequest struct {
	Type          string `json:"type"`
	SourceURL     string `json:"source_url"`
	OriginalTitle string `json:"original_title"`
	Status        string `json:"status"`
}