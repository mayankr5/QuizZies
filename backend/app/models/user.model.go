package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID   `json:"id" gorm:"type:uuid;primaryKey"`
	Name     string      `json:"name"`
	Email    string      `json:"email" gorm:"unique;not null"`
	Username string      `json:"username" gorm:"unique;not null"`
	Password string      `json:"password"`
	Quizzes  []Quiz      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Tokens   []AuthToken `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

type AuthToken struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Token  string    `gorm:"unique;not null" json:"token"`
	UserID uuid.UUID `gorm:"not null" json:"user_id"`
}
