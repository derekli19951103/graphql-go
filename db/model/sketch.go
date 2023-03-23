package model

import (
	"time"

	"gorm.io/gorm"
)

// Sketch represents a sketch created by a user.
type Sketch struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Content   string `gorm:"type:text"`
	User      User
	UserID    int
}