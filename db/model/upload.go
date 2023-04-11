package model

import (
	"gorm.io/gorm"
)

type Upload struct {
    gorm.Model
	ID        int `gorm:"primaryKey"`
    Type      int `gorm:"not null"`
	URL       string `gorm:"not null"`
	User      User
	UserID    int
}

type FileType int
const (
    JSON FileType = 1
    Image  FileType = 2
    Other FileType = 100
)