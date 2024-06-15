package models

import (
	"time"
)

type Project struct {
	UUID      string `gorm:"type:char(36);primaryKey"`
	Name      string `gorm:"not null"`
	Owner     string `gorm:"type:char(36)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
