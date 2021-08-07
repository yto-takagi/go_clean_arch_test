package sql

import (
	"go_clean_arch_test/app/domain"
	"log"

	"github.com/jinzhu/gorm"
)

func GetAll(db *gorm.DB, article []domain.Article) []domain.Article {
	// db.Order("created_at asc").Find(&articles)
	db.Debug().Table("article").Find(&article)
	// log
	log.Println("++++++++++++++++++++++ article_sql.go ++++++++++++++++++++++")
	log.Println(article)

	return article
}
