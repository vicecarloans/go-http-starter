package books

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type bookService struct {
	bookRepository BookRepository
}

func NewBookService(bookRepository BookRepository) BookService {
	return &bookService{bookRepository: bookRepository}
}

func (s *bookService) CreateBook(ctx context.Context, book *CreateBookDTO) error {
	book.ID = uuid.New().String()
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	return s.bookRepository.CreateBook(ctx, book)
}

func (s *bookService) GetBook(ctx context.Context, id string) (*Books, error) {
	return s.bookRepository.GetBook(ctx, id)
}

func (s *bookService) UpdateBook(ctx context.Context, book *UpdateBookDTO) error {
	book.UpdatedAt = time.Now()
	return s.bookRepository.UpdateBook(ctx, book)
}

func (s *bookService) DeleteBook(ctx context.Context, id string) error {
	return s.bookRepository.DeleteBook(ctx, id)
}

func (s *bookService) GetBooks(ctx context.Context, offset int, limit int) ([]Books, error) {
	return s.bookRepository.GetBooks(ctx, offset, limit)
}
