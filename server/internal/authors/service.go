package authors

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type authorService struct {
	authorRepository AuthorRepository
}

func NewAuthorService(authorRepository AuthorRepository) AuthorService {
	return &authorService{authorRepository: authorRepository}
}

func (s *authorService) CreateAuthor(ctx context.Context, author *CreateAuthorDTO) error {
	author.ID = uuid.New().String()
	author.CreatedAt = time.Now()
	author.UpdatedAt = time.Now()
	return s.authorRepository.CreateAuthor(ctx, author)
}

func (s *authorService) GetAuthor(ctx context.Context, id string) (*GetAuthorDTO, error) {
	author, err := s.authorRepository.GetAuthor(ctx, id)
	if err != nil {
		return nil, err
	}

	var books []*GetBookDTO
	for _, book := range author.Books {
		bookDTO := &GetBookDTO{
			BaseBookDTO: BaseBookDTO{
				Title: book.Title,
				Price: book.Price,
			},
			ID:        book.ID,
			CreatedAt: book.CreatedAt,
			UpdatedAt: book.UpdatedAt,
		}
		books = append(books, bookDTO)
	}
	return &GetAuthorDTO{
		BaseAuthorDTO: BaseAuthorDTO{
			FirstName: author.FirstName,
			LastName:  author.LastName,
		},
		ID:        author.BaseAuthor.ID,
		CreatedAt: author.BaseAuthor.CreatedAt,
		UpdatedAt: author.BaseAuthor.UpdatedAt,
		Books:     books,
	}, nil
}

func (s *authorService) UpdateAuthor(ctx context.Context, author *UpdateAuthorDTO) error {
	return s.authorRepository.UpdateAuthor(ctx, author)
}

func (s *authorService) DeleteAuthor(ctx context.Context, id string) error {
	return s.authorRepository.DeleteAuthor(ctx, id)
}

func (s *authorService) GetAuthors(ctx context.Context, offset int, limit int) ([]*GetManyAuthorsDTO, error) {
	authors, err := s.authorRepository.GetAuthors(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	var authorsDTO []*GetManyAuthorsDTO 
	for _, author := range authors {
		authorsDTO = append(authorsDTO, &GetManyAuthorsDTO{
			BaseAuthorDTO: BaseAuthorDTO{
				FirstName: author.FirstName,
				LastName:  author.LastName,
			},
			ID:        author.BaseAuthor.ID,
			CreatedAt: author.BaseAuthor.CreatedAt,
			UpdatedAt: author.BaseAuthor.UpdatedAt,
		})
	}

	return authorsDTO, nil
}
