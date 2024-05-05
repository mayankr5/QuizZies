package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uuid.UUID    `json:"id"`
	FirstName    string       `json:"first_name" validate:"required,min=2,max=100"`
	LastName     string       `json:"last_name" validate:"required,min=2,max=100"`
	Username     string       `json:"username" validate:"required, min=6,max=100"`
	Password     string       `json:"Password" validate:"required,min=6"`
	Email        string       `json:"email" validate:"email,required"`
	Phone        string       `json:"phone" validate:"required"`
	AccessToken  *string      `json:"access_token"`
	RefreshToken *string      `json:"refresh_token"`
	Quizzes      []Quiz       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	ScoreBoards  []ScoreBoard `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
