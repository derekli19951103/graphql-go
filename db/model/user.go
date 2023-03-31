package model

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user of the application.
type User struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
}

// Session represents a user's session on the application.
type Session struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	UserID    int
	User	  User
	Token     string `gorm:"unique;not null"`
	ExpiresAt time.Time
}