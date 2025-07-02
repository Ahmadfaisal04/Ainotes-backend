package dto

import "github.com/google/uuid"

type QuizResponse struct {
	ID                uuid.UUID `json:"id"`
	NoteID            uuid.UUID `json:"note_id"`
	Question          string    `json:"question"`
	Options           string    `json:"options"` // JSON string
	CorrectAnswerIndex int       `json:"correct_answer_index"`
	Explanation       string    `json:"explanation"`
}