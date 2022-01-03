package usecase

import (
	"context"
	"go_clean_arch_test/app/domain"
	channel "go_clean_arch_test/app/domain/channel"
	form "go_clean_arch_test/app/domain/form"
	"go_clean_arch_test/app/interfaces/database/repository/entity"
	"go_clean_arch_test/app/transaction"
	"go_clean_arch_test/app/usecase/repository"
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

	inputChanel(ctx context.Context, article *domain.Article, ch chan *channel.ArticleInputChannel)
	updateExpChanel(ctx context.Context, article *domain.Article, ch chan *channel.ArticleInputChannel)
}

type articleUsecase struct {
	authorUsecase     AuthorUsecase
	expPoolUsecase    ExpPoolUsecase
	lvUsecase         LvUsecase
	articleRepository repository.ArticleRepository
	trancaction       transaction.Transaction
}

// NewArticleUsecase constructor
func NewArticleUsecase(authorUsecase AuthorUsecase, expPoolUsecase ExpPoolUsecase, lvUsecase LvUsecase, articleRepository repository.ArticleRepository, trancaction transaction.Transaction) ArticleUsecase {
	return &articleUsecase{authorUsecase: authorUsecase, expPoolUsecase: expPoolUsecase, lvUsecase: lvUsecase, articleRepository: articleRepository, trancaction: trancaction}
}

// 全件取得
func (articleUsecase *articleUsecase) GetAll(userId int) ([]domain.Article, error) {

	var article []entity.Article
	articles, err := articleUsecase.articleRepository.GetAll(article, userId)
	if err != nil {
		return nil, err
	}

	var articleModelSlice []domain.Article
	for _, article := range articles {
		authorModel, err := domain.NewAuthor(article.Author.Id, article.Author.Name, article.Author.UserId, article.Author.UpdatedAt, article.Author.CreatedAt)
		if err != nil {
			return nil, err
		}
		expPoolModel, err := domain.NewExpPool(article.ExpPool.Id, article.ExpPool.UserId, article.ExpPool.Exp, article.ExpPool.Lv, article.ExpPool.UpdatedAt, article.ExpPool.CreatedAt)
		if err != nil {
			return nil, err
		}
		articleModel, err := domain.NewArticle(article.Id, article.Title, article.Content, article.UpdatedAt, article.CreatedAt, *authorModel, *expPoolModel)
		if err != nil {
			return nil, err
		}
		articleModelSlice = append(articleModelSlice, *articleModel)
	}

	return articleModelSlice, nil
}

// Id指定
func (articleUsecase *articleUsecase) GetById(id int) (domain.Article, error) {

	var article entity.Article
	var articleModel domain.Article
	article, err := articleUsecase.articleRepository.GetById(article, id)
	if err != nil {
		return articleModel, err
	}

	authorModel, err := domain.NewAuthor(article.Author.Id, article.Author.Name, article.Author.UserId, article.Author.UpdatedAt, article.Author.CreatedAt)
	if err != nil {
		return articleModel, err
	}
	expPoolModel, err := domain.NewExpPool(article.ExpPool.Id, article.ExpPool.UserId, article.ExpPool.Exp, article.ExpPool.Lv, article.ExpPool.UpdatedAt, article.ExpPool.CreatedAt)
	if err != nil {
		return articleModel, err
	}
	articleModel.Set(article.Id, article.Title, article.Content, article.UpdatedAt, article.CreatedAt, *authorModel, *expPoolModel)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("GetById",
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return articleModel, nil
}

// Id、ユーザーID指定
func (articleUsecase *articleUsecase) GetByIdAndUserId(id int, userId int) (domain.Article, error) {

	var article entity.Article
	var articleModel domain.Article
	article, err := articleUsecase.articleRepository.GetByIdAndUserId(article, id, userId)
	if err != nil {
		return articleModel, err
	}

	authorModel, err := domain.NewAuthor(article.Author.Id, article.Author.Name, article.Author.UserId, article.Author.UpdatedAt, article.Author.CreatedAt)
	if err != nil {
		return articleModel, err
	}
	expPoolModel, err := domain.NewExpPool(article.ExpPool.Id, article.ExpPool.UserId, article.ExpPool.Exp, article.ExpPool.Lv, article.ExpPool.UpdatedAt, article.ExpPool.CreatedAt)
	if err != nil {
		return articleModel, err
	}
	articleModel.Set(article.Id, article.Title, article.Content, article.UpdatedAt, article.CreatedAt, *authorModel, *expPoolModel)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("GetByIdAndUserId",
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return articleModel, nil
}

// authorId、ユーザーID指定
func (articleUsecase *articleUsecase) GetByAuthorIdAndUserId(id int, userId int) ([]domain.Article, error) {

	var article []entity.Article
	articles, err := articleUsecase.articleRepository.GetByAuthorIdAndUserId(article, id, userId)
	if err != nil {
		return nil, err
	}

	var articleModelSlice []domain.Article
	for _, article := range articles {
		authorModel, err := domain.NewAuthor(article.Author.Id, article.Author.Name, article.Author.UserId, article.Author.UpdatedAt, article.Author.CreatedAt)
		if err != nil {
			return nil, err
		}
		expPoolModel, err := domain.NewExpPool(article.ExpPool.Id, article.ExpPool.UserId, article.ExpPool.Exp, article.ExpPool.Lv, article.ExpPool.UpdatedAt, article.ExpPool.CreatedAt)
		if err != nil {
			return nil, err
		}
		articleModel, err := domain.NewArticle(article.Id, article.Title, article.Content, article.UpdatedAt, article.CreatedAt, *authorModel, *expPoolModel)
		if err != nil {
			return nil, err
		}
		articleModelSlice = append(articleModelSlice, *articleModel)
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("GetByAuthorIdAndUserId",
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return articleModelSlice, nil
}

// 検索
func (articleUsecase *articleUsecase) GetLikeByTitleAndContent(searchContent string, userId int) ([]domain.Article, error) {

	var article []entity.Article
	articles, err := articleUsecase.articleRepository.SearchContent(article, searchContent, userId)
	if err != nil {
		return nil, err
	}

	var articleModelSlice []domain.Article
	for _, article := range articles {
		authorModel, err := domain.NewAuthor(article.Author.Id, article.Author.Name, article.Author.UserId, article.Author.UpdatedAt, article.Author.CreatedAt)
		if err != nil {
			return nil, err
		}
		expPoolModel, err := domain.NewExpPool(article.ExpPool.Id, article.ExpPool.UserId, article.ExpPool.Exp, article.ExpPool.Lv, article.ExpPool.UpdatedAt, article.ExpPool.CreatedAt)
		if err != nil {
			return nil, err
		}
		articleModel, err := domain.NewArticle(article.Id, article.Title, article.Content, article.UpdatedAt, article.CreatedAt, *authorModel, *expPoolModel)
		if err != nil {
			return nil, err
		}
		articleModelSlice = append(articleModelSlice, *articleModel)
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("GetLikeByTitleAndContent",
		zap.String("param searchContent", searchContent),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return articleModelSlice, nil
}

// 新規登録
// トランザクション
func (articleUsecase *articleUsecase) Input(ctx context.Context, article *domain.Article) error {

	_, err := articleUsecase.trancaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {

		ch := make(chan *channel.ArticleInputChannel)
		// 新規登録
		go articleUsecase.inputChanel(ctx, article, ch)
		// 経験値登録、レベルアップ
		go articleUsecase.updateExpChanel(ctx, article, ch)

		article := <-ch
		article = <-ch

		return article, nil
	})

	return err
}

// 更新
// トランザクション
func (articleUsecase *articleUsecase) Update(ctx context.Context, article *domain.Article) error {

	_, err := articleUsecase.trancaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		authorByName, err := articleUsecase.authorUsecase.GetByName(article.Author.GetName(), article.Author.GetUserId())
		if err != nil {
			return article, err
		}
		if authorByName.GetId() == 0 {
			articleUsecase.authorUsecase.Input(ctx, &article.Author)
		} else {
			articleUsecase.authorUsecase.Update(ctx, &article.Author)
		}

		ArticleForm, err := form.NewArticleForm(article.GetId(), article.GetTitle(), article.GetContent(), time.Now(), time.Time{}, article.Author.GetId())
		if err != nil {
			return article, nil
		}

		err = articleUsecase.articleRepository.Update(ctx, ArticleForm)
		if err != nil {
			return article, err
		}

		var expPool domain.ExpPool
		article.Set(ArticleForm.Id, ArticleForm.Title, ArticleForm.Content, ArticleForm.UpdatedAt, ArticleForm.CreatedAt, authorByName, expPool)

		return article, nil
	})

	return err
}

// 削除
func (articleUsecase *articleUsecase) Delete(id int) error {

	articleForm, _ := form.NewArticleForm(id, "delete", "", time.Time{}, time.Time{}, 0)
	err := articleUsecase.articleRepository.Delete(articleForm)
	if err != nil {
		return err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("Delete",
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

		articleForm, _ := form.NewArticleForm(0, "delete", "", time.Time{}, time.Time{}, author.GetId())
		err = articleUsecase.articleRepository.DeleteByAuthorId(ctx, articleForm)
		if err != nil {
			return author, err
		}

		return author, nil
	})

	return err
}

// 新規登録チャネル
func (articleUsecase *articleUsecase) inputChanel(ctx context.Context, article *domain.Article, ch chan *channel.ArticleInputChannel) {

	var err error
	articleInputChannel, err := channel.NewArticleInputChannel(*article, err)
	if err != nil {
		articleInputChannel.Set(*article, err)
		ch <- articleInputChannel
	}

	authorByName, err := articleUsecase.authorUsecase.GetByName(article.Author.GetName(), article.Author.GetUserId())

	if err != nil {
		articleInputChannel.Set(*article, err)
		ch <- articleInputChannel
	}

	if authorByName.GetId() == 0 {
		// カテゴリー存在しなければ、カテゴリー新規登録してそのIdで記事更新
		authorByName, err = articleUsecase.authorUsecase.Input(ctx, &article.Author)
		if err != nil {
			articleInputChannel.Set(*article, err)
			ch <- articleInputChannel
		}
	}

	article.Author.Set(authorByName.GetId(), authorByName.GetName(), authorByName.GetUserId(), authorByName.GetUpdatedAt(), authorByName.GetCreatedAt())
	article.Set(article.GetId(), article.GetTitle(), article.GetContent(), article.GetUpdatedAt(), article.GetCreatedAt(), article.GetAuthor(), article.GetExpPool())

	articleForm, err := form.NewArticleForm(0, article.GetTitle(), article.GetContent(), time.Now(), time.Now(), article.Author.GetId())
	if err != nil {
		articleInputChannel.Set(*article, err)
		ch <- articleInputChannel
	}

	err = articleUsecase.articleRepository.Input(ctx, articleForm)
	if err != nil {
		articleInputChannel.Set(*article, err)
		ch <- articleInputChannel
	}

	article.Set(articleForm.Id, articleForm.Title, articleForm.Content, articleForm.UpdatedAt, articleForm.CreatedAt, article.GetAuthor(), article.GetExpPool())

	articleInputChannel.Set(*article, err)
	ch <- articleInputChannel
}

// 経験値更新・レベルアップチャネル
func (articleUsecase *articleUsecase) updateExpChanel(ctx context.Context, article *domain.Article, ch chan *channel.ArticleInputChannel) {

	var err error
	articleInputChannel, err := channel.NewArticleInputChannel(*article, err)
	if err != nil {
		articleInputChannel.Set(*article, err)
		ch <- articleInputChannel
	}

	// 経験値取得
	expPoolGetByUserId, err := articleUsecase.expPoolUsecase.GetByUserId(article.Author.GetUserId())
	if err != nil {
		articleInputChannel.Set(*article, err)
		ch <- articleInputChannel
	}

	if expPoolGetByUserId.GetId() == 0 {
		// 経験値新規登録
		expPoolGetByUserId.Set(expPoolGetByUserId.GetId(), article.Author.GetUserId(), 1, 1, expPoolGetByUserId.GetUpdatedAt(), time.Now())
		err = articleUsecase.expPoolUsecase.Input(ctx, &expPoolGetByUserId)

		article.ExpPool.Set(expPoolGetByUserId.GetId(), expPoolGetByUserId.GetUserId(), expPoolGetByUserId.GetExp(), expPoolGetByUserId.GetLv(), expPoolGetByUserId.GetUpdatedAt(), expPoolGetByUserId.GetCreatedAt())

		articleInputChannel.Set(*article, err)
		ch <- articleInputChannel
	}

	// レベルアップ判定
	expPoolGetByUserId.Set(expPoolGetByUserId.GetId(), expPoolGetByUserId.GetUserId(), expPoolGetByUserId.GetExp()+1, expPoolGetByUserId.GetLv(), expPoolGetByUserId.GetUpdatedAt(), expPoolGetByUserId.GetCreatedAt())
	lv, err := articleUsecase.lvUsecase.GetByExp(expPoolGetByUserId.GetExp())
	if err != nil {
		articleInputChannel.Set(*article, err)
		ch <- articleInputChannel
	}

	if lv.GetLv() > expPoolGetByUserId.GetLv() {
		// レベルアップ
		expPoolGetByUserId.Set(expPoolGetByUserId.GetId(), expPoolGetByUserId.GetUserId(), expPoolGetByUserId.GetExp(), lv.GetLv(), expPoolGetByUserId.GetUpdatedAt(), expPoolGetByUserId.GetCreatedAt())
	}

	// 経験値更新
	expPoolGetByUserId.Set(expPoolGetByUserId.GetId(), expPoolGetByUserId.GetUserId(), expPoolGetByUserId.GetExp(), expPoolGetByUserId.GetLv(), time.Now(), expPoolGetByUserId.GetCreatedAt())
	err = articleUsecase.expPoolUsecase.Update(ctx, &expPoolGetByUserId)

	article.ExpPool.Set(expPoolGetByUserId.GetId(), expPoolGetByUserId.GetUserId(), expPoolGetByUserId.GetExp(), expPoolGetByUserId.GetLv(), expPoolGetByUserId.GetUpdatedAt(), expPoolGetByUserId.GetCreatedAt())

	articleInputChannel.Set(*article, err)
	ch <- articleInputChannel
}
