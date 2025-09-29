package models

import (
	"time"
)

type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"-"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	Author    string    `gorm:"type:varchar(255);not null" json:"author"`
	ISBN      string    `gorm:"type:varchar(100);uniqueIndex" json:"isbn"`
}
