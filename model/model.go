package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	GoogleID  string    `gorm:"type:varchar;not null"`
	Email     string    `gorm:"type:varchar;not null"`
	FullName  string    `gorm:"type:varchar"`
	AvatarURL string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"type:timestamptz"`
	UpdatedAt time.Time `gorm:"type:timestamptz"`
}

type Resource struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID        uuid.UUID `gorm:"type:uuid;foreignKey:UserID;references:ID"`
	Type          string    `gorm:"type:varchar"`
	SourceURL     string    `gorm:"type:text"`
	OriginalTitle string    `gorm:"type:varchar"`
	Status        string    `gorm:"type:varchar"`
	CreatedAt     time.Time `gorm:"type:timestamptz"`
	UpdatedAt     time.Time `gorm:"type:timestamptz"`
	User          User
}

type Note struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	ResourceID uuid.UUID `gorm:"type:uuid;foreignKey:ResourceID;references:ID"`
	UserID     uuid.UUID `gorm:"type:uuid;foreignKey:UserID;references:ID"`
	Title      string    `gorm:"type:varchar"`
	Summary    string    `gorm:"type:text"`
	FullText   string    `gorm:"type:text"`
	CreatedAt  time.Time `gorm:"type:timestamptz"`
	UpdatedAt  time.Time `gorm:"type:timestamptz"`
	Resource   Resource
	User       User
}

type Quiz struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey"`
	NoteID            uuid.UUID `gorm:"type:uuid;foreignKey:NoteID;references:ID"`
	Question          string    `gorm:"type:text"`
	Options           string    `gorm:"type:jsonb"`
	CorrectAnswerIndex int       `gorm:"type:int"`
	Explanation       string    `gorm:"type:text"`
	Note              Note
}

type Flashcard struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	NoteID    uuid.UUID `gorm:"type:uuid;foreignKey:NoteID;references:ID"`
	FrontText string    `gorm:"type:text"`
	BackText  string    `gorm:"type:text"`
	Note      Note
}