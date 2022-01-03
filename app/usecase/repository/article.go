package repository

import (
	"context"
	form "go_clean_arch_test/app/domain/form"
	"go_clean_arch_test/app/interfaces/database/repository/entity"
)

// AritcleRepository interface
type ArticleRepository interface {
	GetAll(article []entity.Article, userId int) ([]entity.Article, error)
	GetById(article entity.Article, id int) (entity.Article, error)
	GetByIdAndUserId(article entity.Article, id int, userId int) (entity.Article, error)
	GetByAuthorIdAndUserId(articles []entity.Article, id int, userId int) ([]entity.Article, error)
	SearchContent(articles []entity.Article, searchContent string, userId int) ([]entity.Article, error)
	Input(ctx context.Context, articleForm *form.ArticleForm) error
	Update(ctx context.Context, articleForm *form.ArticleForm) error
	Delete(articleForm *form.ArticleForm) error
	DeleteByAuthorId(ctx context.Context, articleForm *form.ArticleForm) error
}
