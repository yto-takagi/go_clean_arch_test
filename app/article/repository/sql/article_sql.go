package sql

import (
	"context"
	"errors"
	"go_clean_arch_test/app/article/database"
	"go_clean_arch_test/app/domain"
	form "go_clean_arch_test/app/domain/form"
	"go_clean_arch_test/app/domain/repository"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// ArticleRepository struct
type ArticleRepository struct {
	Conn *gorm.DB
}

// NewArticleRepository constructor
func NewArticleRepository(conn *gorm.DB) repository.ArticleRepository {
	return &ArticleRepository{Conn: conn}
}

// 全件取得
func (articleRepository *ArticleRepository) GetAll(article []domain.Article, userId int) ([]domain.Article, error) {
	if err := articleRepository.Conn.
		Debug().
		Table("articles").
		Select("articles.id , articles.title, articles.content, articles.created_at, articles.updated_at, articles.deleted_at, articles.author_id, authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Joins("INNER JOIN authors ON articles.author_id = authors.id").
		Where("authors.user_id = ?", userId).
		Scan(&article).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return article, nil
}

// Idに紐付いたデータ取得
func (articleRepository *ArticleRepository) GetById(article domain.Article, id int) (domain.Article, error) {
	if err := articleRepository.Conn.
		Debug().
		Table("articles").
		Select("articles.id , articles.title, articles.content, articles.created_at, articles.updated_at, articles.deleted_at, articles.author_id, authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Joins("INNER JOIN authors ON articles.author_id = authors.id").
		Where("articles.id = ?", id).
		Scan(&article).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return article, err
		}
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("GetById",
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return article, nil

}

// Id、ユーザーIdに紐付いたデータ取得
func (articleRepository *ArticleRepository) GetByIdAndUserId(article domain.Article, id int, userId int) (domain.Article, error) {
	if err := articleRepository.Conn.
		Debug().
		Table("articles").
		Select("articles.id , articles.title, articles.content, articles.created_at, articles.updated_at, articles.deleted_at, articles.author_id, authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Joins("INNER JOIN authors ON articles.author_id = authors.id").
		Where("articles.id = ? AND authors.user_id = ?", id, userId).
		Scan(&article).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return article, err
		}
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("GetByIdAndUserId",
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return article, nil

}

// authorId、ユーザーIdに紐付いたデータ取得
func (articleRepository *ArticleRepository) GetByAuthorIdAndUserId(articles []domain.Article, id int, userId int) ([]domain.Article, error) {
	if err := articleRepository.Conn.
		Debug().
		Table("articles").
		Select("articles.id , articles.title, articles.content, articles.created_at, articles.updated_at, articles.deleted_at, articles.author_id, authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Joins("INNER JOIN authors ON articles.author_id = authors.id").
		Where("authors.id = ? AND authors.user_id = ?", id, userId).
		Scan(&articles).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("GetByAuthorIdAndUserId",
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return articles, nil

}

// titleとcontentを曖昧検索
func (articleRepository *ArticleRepository) SearchContent(articles []domain.Article, searchContent string, userId int) ([]domain.Article, error) {
	if err := articleRepository.Conn.
		Debug().
		Table("articles").
		Select("articles.id , articles.title, articles.content, articles.created_at, articles.updated_at, articles.deleted_at, articles.author_id, authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Joins("INNER JOIN authors ON articles.author_id = authors.id").
		Where("authors.user_id = ? AND (articles.title LIKE ? OR articles.content LIKE ?)",
			userId,
			"%"+searchContent+"%",
			"%"+searchContent+"%",
		).
		Scan(&articles).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("SearchContent",
		zap.String("param searchContent", searchContent),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return articles, nil

}

// 新規登録
func (articleRepository *ArticleRepository) Input(ctx context.Context, articleForm *form.ArticleForm) error {
	dao, ok := database.GetTx(ctx)
	if !ok {
		dao = articleRepository.Conn
	}
	if err := dao.
		Debug().
		Table("articles").
		Create(&articleForm).
		Error; err != nil {
		return err
	}

	return nil

}

// 更新
func (articleRepository *ArticleRepository) Update(ctx context.Context, articleForm *form.ArticleForm) error {
	dao, ok := database.GetTx(ctx)
	if !ok {
		dao = articleRepository.Conn
	}
	if err := dao.
		Debug().
		Model(&articleForm).
		Table("articles").
		Omit("createdAt").
		Where("id = ?", articleForm.Id).
		Updates(map[string]interface{}{
			"title":      articleForm.Title,
			"content":    articleForm.Content,
			"updated_at": articleForm.UpdatedAt,
			"author_id":  articleForm.AuthorId}).
		Error; err != nil {
		return err
	}

	return nil
}

// 削除
func (articleRepository *ArticleRepository) Delete(articleForm *form.ArticleForm) error {
	if err := articleRepository.Conn.
		Debug().
		Table("articles").
		Where("id = ?", articleForm.Id).
		Delete(&articleForm).
		Error; err != nil {
		return err
	}

	return nil
}

// 削除(authorId指定)
func (articleRepository *ArticleRepository) DeleteByAuthorId(ctx context.Context, articleForm *form.ArticleForm) error {
	dao, ok := database.GetTx(ctx)
	if !ok {
		dao = articleRepository.Conn
	}
	if err := dao.
		Debug().
		Table("articles").
		Where("author_id = ?", articleForm.AuthorId).
		Delete(&articleForm).
		Error; err != nil {
		return err
	}

	return nil
}
