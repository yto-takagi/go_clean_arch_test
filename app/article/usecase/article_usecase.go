package usecase

import (
	"go_clean_arch_test/app/article/repository/sql"
	"go_clean_arch_test/app/domain"
	"log"
	"time"

	"go.uber.org/zap"
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
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetAll"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(articles)

	return articles
}

// Id指定
func (usecase *ArticleUsecase) GetById(id int) domain.Article {
	db := usecase.DB.Connect()
	// defer db.Close()

	var article domain.Article
	articleByid := sql.GetById(db, article, id)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetById"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(articleByid)

	return articleByid
}
