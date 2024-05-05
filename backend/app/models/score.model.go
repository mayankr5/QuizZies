package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScoreBoard struct {
	gorm.Model
	ID     uuid.UUID `json:"id" gorm:"primaryKey"`
	QuizID uuid.UUID `json:"quiz_id" gorm:"not null"`
	UserID uuid.UUID `json:"user_id" gorm:"not null"`
	Score  int       `json:"score" gorm:"not null"`
}
