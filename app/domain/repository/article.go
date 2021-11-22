package repository

import (
	"go_clean_arch_test/app/domain"
	form "go_clean_arch_test/app/domain/form"
)

// AritcleRepository interface
type ArticleRepository interface {
	GetAll(article []domain.Article, userId int) ([]domain.Article, error)
	GetById(article domain.Article, id int) (domain.Article, error)
	GetByIdAndUserId(article domain.Article, id int, userId int) (domain.Article, error)
	GetByAuthorIdAndUserId(articles []domain.Article, id int, userId int) ([]domain.Article, error)
	SearchContent(articles []domain.Article, searchContent string, userId int) ([]domain.Article, error)
	Input(articleForm *form.ArticleForm) error
	Update(articleForm *form.ArticleForm) error
	Delete(articleForm *form.ArticleForm) error
	DeleteByAuthorId(articleForm *form.ArticleForm) error
}
