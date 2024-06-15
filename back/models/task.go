package models

import (
	"time"
)

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	ProjectID   string `gorm:"type:char(36);index"`
	Title       string `gorm:"not null"`
	Description string
	Status      string
	Priority    string
	Assignee    string `gorm:"type:char(36)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
