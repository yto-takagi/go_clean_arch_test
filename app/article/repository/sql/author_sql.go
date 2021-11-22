package sql

import (
	"errors"
	"go_clean_arch_test/app/domain"
	"go_clean_arch_test/app/domain/repository"
	"log"
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
func (authorRepository *AuthorRepository) GetAuthorByUser(author []domain.Author, userId int) ([]domain.Author, error) {
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

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_sql.go ++++++++++++++++++++++",
		zap.String("method", "GetAuthorByUser"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(author)

	return author, nil

}

// AuthorId、ユーザーIdに紐付いたデータ取得
func (authorRepository *AuthorRepository) GetAuthorByAuthorIdAndUserId(author domain.Author, id int, userId int) (domain.Author, error) {
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
	logger.Info("++++++++++++++++++++++ author_sql.go ++++++++++++++++++++++",
		zap.String("method", "GetByAuthorIdAndUserId"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(author)

	return author, nil

}

// カテゴリ名に紐付いたデータ取得
func (authorRepository *AuthorRepository) GetByAuthorName(author domain.Author, name string, userId int) (domain.Author, error) {
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
	logger.Info("++++++++++++++++++++++ author_sql.go ++++++++++++++++++++++",
		zap.String("method", "GetByName"),
		zap.String("param name", name),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(author)

	return author, nil

}

// 新規登録
func (authorRepository *AuthorRepository) InputByAuthor(author *domain.Author) error {
	if err := authorRepository.Conn.
		Debug().
		Create(&author).
		Error; err != nil {
		return err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_sql.go ++++++++++++++++++++++",
		zap.String("method", "Input"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(author)

	return nil
}

// 更新
func (authorRepository *AuthorRepository) UpdateByAuthor(author *domain.Author) error {

	if err := authorRepository.Conn.
		Debug().
		Model(&author).Omit("createdAt").
		Updates(map[string]interface{}{"name": author.Name, "updated_at": author.UpdatedAt}).
		Error; err != nil {
		return err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_sql.go ++++++++++++++++++++++",
		zap.String("method", "Update"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(author)

	return nil
}

// 削除
func (authorRepository *AuthorRepository) DeleteByAuthor(author *domain.Author, userId int) error {
	if err := authorRepository.Conn.
		Debug().
		Table("authors").
		Where("id = ? AND user_id = ?", author.Id, userId).
		Delete(&author).
		Error; err != nil {
		return err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_sql.go ++++++++++++++++++++++",
		zap.String("method", "Delete"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return nil
}
