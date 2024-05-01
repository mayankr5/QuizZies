package models

import (
	"github.com/google/uuid"
)

type ScoreBoard struct {
	ID     uuid.UUID `json:"id" gorm:"primaryKey"`
	QuizID uuid.UUID `json:"quiz_id"`
	User   User      `json:"user"`
	Score  int       `json:"score"`
}
