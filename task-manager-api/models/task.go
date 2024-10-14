package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Description string         `gorm:"type:varchar(255);not null" json:"description"`
	Status      string         `gorm:"type:varchar(255);not null" json:"status"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
