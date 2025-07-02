package dto

import "github.com/google/uuid"

type QuizCreateRequest struct {
	NoteID            uuid.UUID `json:"note_id" binding:"required"`
	Question          string    `json:"question" binding:"required"`
	Options           string    `json:"options" binding:"required"` // JSON string
	CorrectAnswerIndex int       `json:"correct_answer_index" binding:"required"`
	Explanation       string    `json:"explanation"`
}

type QuizUpdateRequest struct {
	Question          string `json:"question"`
	Options           string `json:"options"` // JSON string
	CorrectAnswerIndex int    `json:"correct_answer_index"`
	Explanation       string `json:"explanation"`
}