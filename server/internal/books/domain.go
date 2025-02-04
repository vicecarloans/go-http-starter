package books

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model

	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	AuthorID  string `gorm:"type:text;foreignKey:ID;references:ID;constraint:OnDelete:CASCADE"`
	Price     float64
}

type BaseBookDTO struct {
	Title     string    `json:"title"`
	Price     float64   `json:"price"`
} 

type CreateBookDTO struct {
	BaseBookDTO
	ID string
	AuthorID string `json:"author_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateBookDTO struct {
	BaseBookDTO
	ID string
	AuthorID string `json:"author_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetBookDTO struct {
	BaseBookDTO
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookRepository interface {
	CreateBook(ctx context.Context, book *CreateBookDTO) error
	GetBook(ctx context.Context, id string) (*Books, error)
	GetBooks(ctx context.Context, offset int, limit int) ([]Books, error)
	UpdateBook(ctx context.Context, book *UpdateBookDTO) error
	DeleteBook(ctx context.Context, id string) error
}

type BookService interface {
	CreateBook(ctx context.Context, book *CreateBookDTO) error
	GetBook(ctx context.Context, id string) (*Books, error)
	GetBooks(ctx context.Context, offset int, limit int) ([]Books, error)
	UpdateBook(ctx context.Context, book *UpdateBookDTO) error
	DeleteBook(ctx context.Context, id string) error
}