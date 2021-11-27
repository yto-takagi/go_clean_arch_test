package usecase

import (
	"context"
	"go_clean_arch_test/app/article/transaction"
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
	Input(ctx context.Context, article *domain.Article) error
	Update(ctx context.Context, article *domain.Article) error
	Delete(id int) error
	DeleteByAuthor(ctx context.Context, author *domain.Author, userId int) error
}
type articleUsecase struct {
	authorUsecase     AuthorUsecase
	articleRepository repository.ArticleRepository
	trancaction       transaction.Transaction
}

// NewArticleUsecase constructor
func NewArticleUsecase(authorUsecase AuthorUsecase, articleRepository repository.ArticleRepository, trancaction transaction.Transaction) ArticleUsecase {
	return &articleUsecase{authorUsecase: authorUsecase, articleRepository: articleRepository, trancaction: trancaction}
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
// トランザクション
func (articleUsecase *articleUsecase) Input(ctx context.Context, article *domain.Article) error {

	_, err := articleUsecase.trancaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {

		// カテゴリー検索(カテゴリー名で)
		authorByName, err := articleUsecase.authorUsecase.GetByName(article.Author.Name, article.Author.UserId)
		if err != nil {
			return article, err
		}
		// if err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		// 	return
		// }
		// TODO 空チェックできてる？
		if authorByName.Id == 0 {
			// カテゴリー存在しなければ、カテゴリー新規登録してそのIdで記事更新
			log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■カテゴリー存在しない場合■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
			authorByName, err = articleUsecase.authorUsecase.Input(ctx, &article.Author)
			if err != nil {
				return article, err
			}
			// if err != nil {
			// 	ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
			// 	return
			// }
		}

		log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■authorByName.Id■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
		log.Println(authorByName.Id)

		article.Author.Id = authorByName.Id

		ArticleForm := form.ArticleForm{}
		ArticleForm.Title = article.Title
		ArticleForm.Content = article.Content
		ArticleForm.CreatedAt = time.Now()
		ArticleForm.UpdatedAt = time.Now()
		ArticleForm.AuthorId = article.Author.Id

		err = articleUsecase.articleRepository.Input(ctx, &ArticleForm)
		if err != nil {
			return article, err
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
		return article, nil
	})

	return err
}

// 更新
// トランザクション
func (articleUsecase *articleUsecase) Update(ctx context.Context, article *domain.Article) error {

	_, err := articleUsecase.trancaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		// TODO カテゴリーが変わってる場合
		// カテゴリー検索(カテゴリー名で)
		authorByName, err := articleUsecase.authorUsecase.GetByName(article.Author.Name, article.Author.UserId)
		if err != nil {
			return article, err
		}
		// TODO 空チェックできてる？
		if authorByName.Id == 0 {
			// カテゴリー存在しなければ、カテゴリー新規登録してそのIdで記事更新
			log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■カテゴリー存在しない場合■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
			articleUsecase.authorUsecase.Input(ctx, &article.Author)
		} else {
			// カテゴリー存在したらそのIdで記事更新
			log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■カテゴリー存在しない場合■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
			articleUsecase.authorUsecase.Update(ctx, &article.Author)
		}

		log.Println("○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○テスト○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○")
		log.Println(&article)

		ArticleForm := form.ArticleForm{}
		ArticleForm.Id = article.Id
		ArticleForm.Title = article.Title
		ArticleForm.Content = article.Content
		ArticleForm.UpdatedAt = time.Now()
		ArticleForm.AuthorId = article.Author.Id

		err = articleUsecase.articleRepository.Update(ctx, &ArticleForm)
		if err != nil {
			return article, err
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

		return article, nil
	})

	return err
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
func (articleUsecase *articleUsecase) DeleteByAuthor(ctx context.Context, author *domain.Author, userId int) error {

	_, err := articleUsecase.trancaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		err := articleUsecase.authorUsecase.Delete(ctx, author, userId)
		if err != nil {
			return author, err
		}

		ArticleForm := form.ArticleForm{}
		ArticleForm.AuthorId = author.Id
		err = articleUsecase.articleRepository.DeleteByAuthorId(ctx, &ArticleForm)
		if err != nil {
			return author, err
		}

		return author, nil
	})

	return err
}
