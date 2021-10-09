package sql

import (
	"go_clean_arch_test/app/domain"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// カテゴリ名に紐付いたデータ取得
func GetByAuthorName(db *gorm.DB, author domain.Author, name string) domain.Author {
	db.Debug().Table("authors").Select("authors.id, authors.name, authors.created_at, authors.updated_at, authors.deleted_at").Where("authors.name = ?", name).Scan(&author)

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
