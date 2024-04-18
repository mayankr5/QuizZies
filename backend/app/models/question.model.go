package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID      uuid.UUID `json:"id" gorm:"primaryKey"`
	Name    string    `json:"name" gorm:"not null"`
	QuizID  uuid.UUID `json:"quiz_id" gorm:"not null"`
	Options []Option  `gorm:"not null;foreignKey:QuestionID;constraint:OnDelete:CASCADE"`
}

type Option struct {
	gorm.Model
	ID         uuid.UUID `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"not null"`
	IsCorrect  bool      `json:"is_correct"`
	QuestionID uuid.UUID `json:"question_id" gorm:"not null"`
}
