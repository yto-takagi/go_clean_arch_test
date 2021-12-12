package sql

import (
	"context"
	"errors"
	"go_clean_arch_test/app/article/database"
	"go_clean_arch_test/app/article/repository/entity"
	"go_clean_arch_test/app/domain/repository"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// AuthorRepository struct
type AuthorRepository struct {
	Conn *gorm.DB
}

// NewAuthorRepository constructor
func NewAuthorRepository(conn *gorm.DB) repository.AuthorRepository {
	return &AuthorRepository{Conn: conn}
}

// ユーザーIDに紐づくデータ取得
func (authorRepository *AuthorRepository) GetAuthorByUser(author []entity.Author, userId int) ([]entity.Author, error) {
	if err := authorRepository.Conn.
		Debug().
		Table("authors").
		Select("authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Where("authors.user_id = ?", userId).
		Scan(&author).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return author, nil

}

// AuthorId、ユーザーIdに紐付いたデータ取得
func (authorRepository *AuthorRepository) GetAuthorByAuthorIdAndUserId(author entity.Author, id int, userId int) (entity.Author, error) {
	if err := authorRepository.Conn.
		Debug().
		Table("authors").
		Select("authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Where("authors.id = ? AND authors.user_id = ?", id, userId).
		Scan(&author).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return author, err
		}
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("GetByAuthorIdAndUserId",
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return author, nil

}

// カテゴリ名に紐付いたデータ取得
func (authorRepository *AuthorRepository) GetByAuthorName(author entity.Author, name string, userId int) (entity.Author, error) {
	if err := authorRepository.Conn.
		Debug().
		Table("authors").
		Select("authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Where("authors.name = ? AND authors.user_id = ?", name, userId).
		Scan(&author).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return author, err
		}
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("GetByName",
		zap.String("param name", name),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return author, nil

}

// 新規登録
func (authorRepository *AuthorRepository) InputByAuthor(ctx context.Context, author *entity.Author) error {
	dao, ok := database.GetTx(ctx)
	if !ok {
		dao = authorRepository.Conn
	}
	if err := dao.
		Debug().
		Create(&author).
		Error; err != nil {
		return err
	}

	return nil
}

// 更新
func (authorRepository *AuthorRepository) UpdateByAuthor(ctx context.Context, author *entity.Author) error {
	dao, ok := database.GetTx(ctx)
	if !ok {
		dao = authorRepository.Conn
	}
	if err := dao.
		Debug().
		Model(&author).Omit("createdAt").
		Updates(map[string]interface{}{"name": author.Name, "updated_at": author.UpdatedAt}).
		Error; err != nil {
		return err
	}

	return nil
}

// 削除
func (authorRepository *AuthorRepository) DeleteByAuthor(ctx context.Context, author *entity.Author, userId int) error {
	dao, ok := database.GetTx(ctx)
	if !ok {
		dao = authorRepository.Conn
	}
	if err := dao.
		Debug().
		Table("authors").
		Where("id = ? AND user_id = ?", author.Id, userId).
		Delete(&author).
		Error; err != nil {
		return err
	}

	return nil
}
