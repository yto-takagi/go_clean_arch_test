package usecase

import (
	"go_clean_arch_test/app/domain"
	form "go_clean_arch_test/app/domain/form"
	"go_clean_arch_test/app/domain/repository"
	"log"
	"time"

	"go.uber.org/zap"
)

// ArticleUsecase interface
type ArticleUsecase interface {
	GetAll(userId int) ([]domain.Article, error)
	GetById(id int) (domain.Article, error)
	GetByIdAndUserId(id int, userId int) (domain.Article, error)
	GetByAuthorIdAndUserId(id int, userId int) ([]domain.Article, error)
	GetLikeByTitleAndContent(searchContent string, userId int) ([]domain.Article, error)
	Input(article *domain.Article) error
	Update(article *domain.Article) error
	Delete(id int) error
	DeleteByAuthor(authorId int) error
}
type articleUsecase struct {
	articleRepository repository.ArticleRepository
}

// NewArticleUsecase constructor
func NewArticleUsecase(articleRepository repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{articleRepository: articleRepository}
}

// 全件取得
func (articleUsecase *articleUsecase) GetAll(userId int) ([]domain.Article, error) {
	// db := usecase.DB.Connect()
	// defer db.Close()

	var article []domain.Article
	articles, err := articleUsecase.articleRepository.GetAll(article, userId)
	if err != nil {
		return nil, err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetAll"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(articles)

	return articles, nil
}

// Id指定
func (articleUsecase *articleUsecase) GetById(id int) (domain.Article, error) {

	var article domain.Article
	articleById, err := articleUsecase.articleRepository.GetById(article, id)
	if err != nil {
		return article, err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetById"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(articleById)

	return articleById, nil
}

// Id、ユーザーID指定
func (articleUsecase *articleUsecase) GetByIdAndUserId(id int, userId int) (domain.Article, error) {

	var article domain.Article
	articleByIdAndUserId, err := articleUsecase.articleRepository.GetByIdAndUserId(article, id, userId)
	if err != nil {
		return article, err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetByIdAndUserId"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(articleByIdAndUserId)

	return articleByIdAndUserId, nil
}

// authorId、ユーザーID指定
func (articleUsecase *articleUsecase) GetByAuthorIdAndUserId(id int, userId int) ([]domain.Article, error) {

	var articles []domain.Article
	articleByIdAndUserId, err := articleUsecase.articleRepository.GetByAuthorIdAndUserId(articles, id, userId)
	if err != nil {
		return nil, err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetByAuthorIdAndUserId"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(articleByIdAndUserId)

	return articleByIdAndUserId, nil
}

// 検索
func (articleUsecase *articleUsecase) GetLikeByTitleAndContent(searchContent string, userId int) ([]domain.Article, error) {

	var article []domain.Article
	articles, err := articleUsecase.articleRepository.SearchContent(article, searchContent, userId)
	if err != nil {
		return nil, err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetLikeByTitleAndContent"),
		zap.String("param searchContent", searchContent),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(articles)

	return articles, nil
}

// 新規登録
func (articleUsecase *articleUsecase) Input(article *domain.Article) error {

	ArticleForm := form.ArticleForm{}
	ArticleForm.Title = article.Title
	ArticleForm.Content = article.Content
	ArticleForm.CreatedAt = time.Now()
	ArticleForm.UpdatedAt = time.Now()
	ArticleForm.AuthorId = article.Author.Id

	err := articleUsecase.articleRepository.Input(&ArticleForm)
	if err != nil {
		return err
	}

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

	return nil
}

// 更新
func (articleUsecase *articleUsecase) Update(article *domain.Article) error {

	ArticleForm := form.ArticleForm{}
	ArticleForm.Id = article.Id
	ArticleForm.Title = article.Title
	ArticleForm.Content = article.Content
	ArticleForm.UpdatedAt = time.Now()
	ArticleForm.AuthorId = article.Author.Id

	err := articleUsecase.articleRepository.Update(&ArticleForm)
	if err != nil {
		return err
	}

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

	return nil
}

// 削除
func (articleUsecase *articleUsecase) Delete(id int) error {

	ArticleForm := form.ArticleForm{}
	ArticleForm.Id = id
	err := articleUsecase.articleRepository.Delete(&ArticleForm)
	if err != nil {
		return err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "Delete"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return nil
}

// 削除(authorId指定)
func (articleUsecase *articleUsecase) DeleteByAuthor(authorId int) error {

	log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■authorId■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
	log.Println(authorId)
	ArticleForm := form.ArticleForm{}
	ArticleForm.AuthorId = authorId
	err := articleUsecase.articleRepository.DeleteByAuthorId(&ArticleForm)
	if err != nil {
		return err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "DeleteByAuthor"),
		zap.Int("param authorId", authorId),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return nil
}
