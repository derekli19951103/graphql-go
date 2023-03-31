package model

import (
	"time"

	"gorm.io/gorm"

	"gorm.io/datatypes"
)

// Sketch represents a sketch created by a user.
type GeoObject struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Type     string
	Content   string `gorm:"type:text"`
	Properties  datatypes.JSON 
	User      User
	UserID    int
}