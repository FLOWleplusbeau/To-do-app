package models

import (
	"time"
)

type User struct {
	ID        string `gorm:"type:char(36);primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsActive  bool `gorm:"default:true"`
}
