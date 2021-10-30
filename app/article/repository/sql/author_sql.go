package sql

import (
	"go_clean_arch_test/app/domain"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// ユーザーIDに紐づくデータ取得
func GetAuthorByUser(db *gorm.DB, author []domain.Author, userId int) []domain.Author {
	db.Debug().Table("authors").
		Select("authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Where("authors.user_id = ?", userId).
		Scan(&author)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_sql.go ++++++++++++++++++++++",
		zap.String("method", "GetAuthorByUser"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(author)

	return author

}

// AuthorId、ユーザーIdに紐付いたデータ取得
func GetAuthorByAuthorIdAndUserId(db *gorm.DB, author domain.Author, id int, userId int) domain.Author {
	db.
		Debug().
		Table("authors").
		Select("authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Where("authors.id = ? AND authors.user_id = ?", id, userId).
		Scan(&author)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_sql.go ++++++++++++++++++++++",
		zap.String("method", "GetByAuthorIdAndUserId"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(author)

	return author

}

// カテゴリ名に紐付いたデータ取得
func GetByAuthorName(db *gorm.DB, author domain.Author, name string, userId int) domain.Author {
	db.Debug().Table("authors").
		Select("authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").
		Where("authors.name = ? AND authors.user_id = ?", name, userId).
		Scan(&author)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_sql.go ++++++++++++++++++++++",
		zap.String("method", "GetByName"),
		zap.String("param name", name),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(author)

	return author

}

// 新規登録
func InputByAuthor(db *gorm.DB, author *domain.Author) {
	db.Debug().Create(&author)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_sql.go ++++++++++++++++++++++",
		zap.String("method", "Input"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(author)

}

// 更新
func UpdateByAuthor(db *gorm.DB, author *domain.Author) {

	// TODO 更新SQL
	db.Debug().Model(&author).Omit("createdAt").Updates(map[string]interface{}{"name": author.Name, "updated_at": author.UpdatedAt})

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_sql.go ++++++++++++++++++++++",
		zap.String("method", "Update"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(author)

}

// 削除
func DeleteByAuthor(db *gorm.DB, author *domain.Author, userId int) {
	db.
		Debug().
		Table("authors").
		Where("id = ? AND user_id = ?", author.Id, userId).
		Delete(&author)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_sql.go ++++++++++++++++++++++",
		zap.String("method", "Delete"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

}
