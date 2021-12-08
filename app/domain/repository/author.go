package repository

import (
	"context"
	"go_clean_arch_test/app/domain"
)

// AuthorRepository interface
type AuthorRepository interface {
	GetAuthorByUser(author []domain.Author, userId int) ([]domain.Author, error)
	GetAuthorByAuthorIdAndUserId(author domain.Author, id int, userId int) (domain.Author, error)
	GetByAuthorName(author domain.Author, name string, userId int) (domain.Author, error)
	InputByAuthor(ctx context.Context, author *domain.Author) error
	UpdateByAuthor(ctx context.Context, author *domain.Author) error
	DeleteByAuthor(ctx context.Context, author *domain.Author, userId int) error
}
