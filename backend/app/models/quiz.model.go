package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	ID          uuid.UUID    `json:"id" gorm:"primaryKey"`
	Title       string       `json:"title" gorm:"not null"`
	Description string       `json:"description"`
	Category    string       `json:"category" gorm:"not null"`
	IsPublic    bool         `json:"is_public" gorm:"default:false"`
	UserID      uuid.UUID    `json:"user_id" gorm:"not null"`
	Questions   []Question   `gorm:"foreignKey:QuizID;constraint:OnDelete:CASCADE"`
	ScoreBoards []ScoreBoard `gorm:"foreignKey:QuizID;constraint:OnDelete:CASCADE;"`
}
