package books

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model

	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Author    string
	Price     float64
}

type BookDTO struct {
	ID        string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Price     float64   `json:"price"`
}
