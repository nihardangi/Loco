package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Amount    float64        `gorm:"not null" json:"amount"`
	Type      string         `gorm:"not null" json:"type"`
	ParentID  uint           `json:"parent_id,omitempty"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
