package usecase

import (
	"go_clean_arch_test/app/article/repository/sql"
	"go_clean_arch_test/app/domain"
	"log"
)

type ArticleUsecase struct {
	DB DBRepository
}

// 全件取得
func (usecase *ArticleUsecase) GetAll() []domain.Article {
	db := usecase.DB.Connect()
	// defer db.Close()

	var article []domain.Article
	articles := sql.GetAll(db, article)
	// log
	log.Println("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++")
	log.Println(articles)
	return articles
}
