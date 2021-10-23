package usecase

import (
	"go_clean_arch_test/app/article/repository/sql"
	"go_clean_arch_test/app/domain"
	form "go_clean_arch_test/app/domain/form"
	"log"
	"time"

	"go.uber.org/zap"
)

type ArticleUsecase struct {
	DB DBRepository
}

// 全件取得
func (usecase *ArticleUsecase) GetAll(userId int) []domain.Article {
	db := usecase.DB.Connect()
	// defer db.Close()

	var article []domain.Article
	articles := sql.GetAll(db, article, userId)

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

// Id、ユーザーID指定
func (usecase *ArticleUsecase) GetByIdAndUserId(id int, userId int) domain.Article {
	db := usecase.DB.Connect()
	// defer db.Close()

	var article domain.Article
	articleByIdAndUserId := sql.GetByIdAndUserId(db, article, id, userId)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetById"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(articleByIdAndUserId)

	return articleByIdAndUserId
}

// 新規登録
func (usecase *ArticleUsecase) Input(article *domain.Article) {
	db := usecase.DB.Connect()
	// defer db.Close()

	ArticleForm := form.ArticleForm{}
	ArticleForm.Title = article.Title
	ArticleForm.Content = article.Content
	ArticleForm.CreatedAt = time.Now()
	ArticleForm.UpdatedAt = time.Now()
	ArticleForm.AuthorId = article.Author.Id

	sql.Input(db, &ArticleForm)

	article.Id = ArticleForm.Id
	article.CreatedAt = ArticleForm.CreatedAt
	article.UpdatedAt = ArticleForm.UpdatedAt

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "Input"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(article)
}

// 更新
func (usecase *ArticleUsecase) Update(article *domain.Article) {
	db := usecase.DB.Connect()
	// defer db.Close()

	ArticleForm := form.ArticleForm{}
	ArticleForm.Id = article.Id
	ArticleForm.Title = article.Title
	ArticleForm.Content = article.Content
	ArticleForm.UpdatedAt = time.Now()
	ArticleForm.AuthorId = article.Author.Id

	// TODO 更新SQL
	sql.Update(db, &ArticleForm)

	article.Id = ArticleForm.Id
	article.CreatedAt = ArticleForm.CreatedAt
	article.UpdatedAt = ArticleForm.UpdatedAt

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "Update"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(article)
}

// 削除
func (usecase *ArticleUsecase) Delete(id int) {
	db := usecase.DB.Connect()
	// defer db.Close()

	ArticleForm := form.ArticleForm{}
	ArticleForm.Id = id
	sql.Delete(db, &ArticleForm)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "Delete"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
}
