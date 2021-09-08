package sql

import (
	"go_clean_arch_test/app/domain"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func GetAll(db *gorm.DB, article []domain.Article) []domain.Article {
	// db.Order("created_at asc").Find(&articles)
	// db.Debug().Table("article").Find(&article)
	db.Debug().Table("article").Select("article.id , article.title, article.content, article.created_at, article.updated_at, article.deleted_at, article.author_id, author.id, author.name, author.created_at, author.updated_at, author.deleted_at").Joins("INNER JOIN author ON article.author_id = author.id").Scan(&article)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_sql.go ++++++++++++++++++++++",
		zap.String("method", "GetAll"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(article)

	return article
}

func GetById(db *gorm.DB, article domain.Article, id int) domain.Article {
	db.Debug().Table("article").Select("article.id , article.title, article.content, article.created_at, article.updated_at, article.deleted_at, article.author_id, author.id, author.name, author.created_at, author.updated_at, author.deleted_at").Joins("INNER JOIN author ON article.author_id = author.id").Where("article.id = ?", id).Scan(&article)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_sql.go ++++++++++++++++++++++",
		zap.String("method", "GetById"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(article)

	return article

}
