package models

import (
	"time"

	"gorm.io/gorm"
)

type UserMovie struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	UserID    uint           `gorm:"not null;index" json:"user_id"`
	MovieID   uint           `gorm:"not null;index" json:"movie_id"`
	Status    string         `gorm:"not null" json:"status"` // "watched" or "want"
	Rating    float32        `gorm:"default:0" json:"rating"`
	Review    string         `json:"review"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Movie     Movie          `json:"movie" gorm:"foreignKey:MovieID"`
}
