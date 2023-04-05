package model

import (
	"time"

	"gorm.io/gorm"

	"gorm.io/datatypes"
)

// GeoObject represents a geo location with data created by a user.
type GeoObject struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Type     string
	Title    string
	Content   string `gorm:"type:text"`
	ImageUrl    string
	Properties  datatypes.JSON 
	User      User
	UserID    int
	Lng 	  float64
	Lat 	  float64
}