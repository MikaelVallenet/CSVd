package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}
