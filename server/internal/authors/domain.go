package authors

import (
	"context"
	"go-http-server/internal/books"
	"time"

	"gorm.io/gorm"
)

type BaseAuthor struct {
	ID        string `gorm:"primaryKey"`
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Authors struct {
	gorm.Model
	BaseAuthor
	Books     []books.Books `gorm:"foreignKey:AuthorID"`
}

type BaseBookDTO = books.BaseBookDTO
type GetBookDTO = books.GetBookDTO

type BaseAuthorDTO struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

type CreateAuthorDTO struct {
	BaseAuthorDTO
	ID        string    
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

type UpdateAuthorDTO struct {
	BaseAuthorDTO
	ID        string    
	CreatedAt time.Time 
	UpdatedAt time.Time 
}

type GetAuthorDTO struct {
	BaseAuthorDTO
	ID        string         `json:"id"`
	Books     []*GetBookDTO `json:"books"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type GetManyAuthorsDTO struct {
	BaseAuthorDTO
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type AuthorRepository interface {
	CreateAuthor(ctx context.Context, author *CreateAuthorDTO) error
	GetAuthor(ctx context.Context, id string) (*Authors, error)
	GetAuthors(ctx context.Context, offset int, limit int) ([]*Authors, error)
	UpdateAuthor(ctx context.Context, author *UpdateAuthorDTO) error
	DeleteAuthor(ctx context.Context, id string) error
}

type AuthorService interface {
	CreateAuthor(ctx context.Context, author *CreateAuthorDTO) error
	GetAuthor(ctx context.Context, id string) (*GetAuthorDTO, error)
	GetAuthors(ctx context.Context, offset int, limit int) ([]*GetManyAuthorsDTO, error)
	UpdateAuthor(ctx context.Context, author *UpdateAuthorDTO) error
	DeleteAuthor(ctx context.Context, id string) error
}

