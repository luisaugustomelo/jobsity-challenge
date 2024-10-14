package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `gorm:"type:varchar(255);not null"`
	Accept      bool   `gorm:"default:false"`
	Status      string `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
