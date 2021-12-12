package repository

import (
	"context"
	"go_clean_arch_test/app/article/repository/entity"
)

// AuthorRepository interface
type AuthorRepository interface {
	GetAuthorByUser(author []entity.Author, userId int) ([]entity.Author, error)
	GetAuthorByAuthorIdAndUserId(author entity.Author, id int, userId int) (entity.Author, error)
	GetByAuthorName(author entity.Author, name string, userId int) (entity.Author, error)
	InputByAuthor(ctx context.Context, author *entity.Author) error
	UpdateByAuthor(ctx context.Context, author *entity.Author) error
	DeleteByAuthor(ctx context.Context, author *entity.Author, userId int) error
}
