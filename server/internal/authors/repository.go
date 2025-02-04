package authors

import (
	"context"

	"gorm.io/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &gormRepository{db: db}
}

func (r *gormRepository) CreateAuthor(ctx context.Context, authorDTO *CreateAuthorDTO) error {
	author := &Authors {
		BaseAuthor: BaseAuthor{
			ID: authorDTO.ID,
			FirstName: authorDTO.FirstName,
			LastName: authorDTO.LastName,
			CreatedAt: authorDTO.CreatedAt,
			UpdatedAt: authorDTO.UpdatedAt,
		},
	}
	return r.db.Model(&Authors{}).Create(author).Error
}

func (r *gormRepository) GetAuthor(ctx context.Context, id string) (*Authors, error) {
	var author Authors
	err := r.db.Model(&Authors{}).
				InnerJoins("Books").
				First(&author, id).
				Error
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *gormRepository) UpdateAuthor(ctx context.Context, updateAuthorDTO *UpdateAuthorDTO) error {
	author := &Authors {
		BaseAuthor: BaseAuthor{
			FirstName: updateAuthorDTO.FirstName,
			LastName: updateAuthorDTO.LastName,
			UpdatedAt: updateAuthorDTO.UpdatedAt,
		},
	}
	return r.db.Model(&Authors{}).Save(author).Error
}

func (r *gormRepository) DeleteAuthor(ctx context.Context, id string) error {
	return r.db.Model(&Authors{}).Delete(id).Error
}

func (r *gormRepository) GetAuthors(ctx context.Context, offset int, limit int) ([]*Authors, error) {
	var authors []*Authors
	err := r.db.Model(&Authors{}).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&authors).Error

	if err != nil {
		return nil, err
	}
	return authors, nil
}