package books

import (
	"context"

	"gorm.io/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &gormRepository{db: db}
}

func (r *gormRepository) GetBook(ctx context.Context, id string) (*Books, error) {
	var book Books
	err := r.db.Model(&Books{}).First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *gormRepository) CreateBook(ctx context.Context, book *CreateBookDTO) error {
	return r.db.Model(&Books{}).Create(book).Error
}

func (r *gormRepository) DeleteBook(ctx context.Context, id string) error {
	return r.db.Model(&Books{}).Delete(&Books{}, id).Error
}

func (r *gormRepository) GetBooks(ctx context.Context, offset int, limit int) ([]Books, error) {
	var books []Books
	err := r.db.Model(&Books{}).
				Order("created_at DESC").
				Offset(offset).
				Limit(limit).
				Find(&books).
				Error
	return books, err
}


func (r *gormRepository) UpdateBook(ctx context.Context, book *UpdateBookDTO) error {
	return r.db.Model(&Books{}).Save(book).Error
}
