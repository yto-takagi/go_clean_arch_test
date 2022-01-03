package usecase

import (
	"context"
	"go_clean_arch_test/app/domain"
	channel "go_clean_arch_test/app/domain/channel"
	"go_clean_arch_test/app/interfaces/database/repository/entity"
	"go_clean_arch_test/app/usecase/repository"
	"time"

	"go.uber.org/zap"
)

// AuthorUsecase interface
type AuthorUsecase interface {
	GetByUser(userId int) ([]domain.Author, domain.ExpPool, error)
	GetByAuthorIdAndUserId(id int, userId int) (domain.Author, error)
	GetByName(name string, userId int) (domain.Author, error)
	Input(ctx context.Context, author *domain.Author) (domain.Author, error)
	Update(ctx context.Context, author *domain.Author) error
	Delete(ctx context.Context, author *domain.Author, userId int) error

	getAuthorByUserChannel(userId int, ch chan *channel.AuthorGetChannel)
	getExpByUserChannel(userId int, ch chan *channel.ExpPoolGetChannel)
}
type authorUsecase struct {
	expPoolUsecase   ExpPoolUsecase
	authorRepository repository.AuthorRepository
}

// NewAuthorUsecase constructor
func NewAuthorUsecase(expPoolUsecase ExpPoolUsecase, authorRepository repository.AuthorRepository) AuthorUsecase {
	return &authorUsecase{expPoolUsecase: expPoolUsecase, authorRepository: authorRepository}
}

// ユーザーIDに紐づくカテゴリー全件取得
func (authorUsecase *authorUsecase) GetByUser(userId int) ([]domain.Author, domain.ExpPool, error) {

	chByGetAuthor := make(chan *channel.AuthorGetChannel)
	chByGetExp := make(chan *channel.ExpPoolGetChannel)
	// 新規登録
	go authorUsecase.getAuthorByUserChannel(userId, chByGetAuthor)
	// 経験値登録、レベルアップ
	go authorUsecase.getExpByUserChannel(userId, chByGetExp)

	author := <-chByGetAuthor
	expPool := <-chByGetExp
	var err error
	if author.GetErr() != nil {
		err = author.GetErr()
	} else if expPool.GetErr() != nil {
		err = expPool.GetErr()
	}

	return author.GetAuthor(), expPool.GetExpPool(), err
}

// Id、ユーザーID指定
func (authorUsecase *authorUsecase) GetByAuthorIdAndUserId(id int, userId int) (domain.Author, error) {

	var author entity.Author
	var authorModel domain.Author
	author, err := authorUsecase.authorRepository.GetAuthorByAuthorIdAndUserId(author, id, userId)
	if err != nil {
		return authorModel, err
	}

	authorModel.Set(author.Id, author.Name, author.UserId, author.UpdatedAt, author.CreatedAt)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("GetById",
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return authorModel, nil
}

// カテゴリー名検索
func (authorUsecase *authorUsecase) GetByName(name string, userId int) (domain.Author, error) {

	var author entity.Author
	var authorModel domain.Author
	author, err := authorUsecase.authorRepository.GetByAuthorName(author, name, userId)
	if err != nil {
		return authorModel, err
	}

	authorModel.Set(author.Id, author.Name, author.UserId, author.UpdatedAt, author.CreatedAt)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("GetByName",
		zap.String("param name", name),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return authorModel, nil
}

// 新規登録
func (authorUsecase *authorUsecase) Input(ctx context.Context, author *domain.Author) (domain.Author, error) {

	var authorEntity entity.Author
	authorEntity.Id = 0
	authorEntity.Name = author.GetName()
	authorEntity.UserId = author.GetUserId()
	authorEntity.UpdatedAt = time.Now()
	authorEntity.CreatedAt = time.Now()

	err := authorUsecase.authorRepository.InputByAuthor(ctx, &authorEntity)
	if err != nil {
		return *author, err
	}

	return *author, nil
}

// 更新
func (authorUsecase *authorUsecase) Update(ctx context.Context, author *domain.Author) error {

	var authorEntity entity.Author
	authorEntity.Id = author.GetId()
	authorEntity.Name = author.GetName()
	authorEntity.UserId = author.GetUserId()
	authorEntity.UpdatedAt = time.Now()
	authorEntity.CreatedAt = author.GetCreatedAt()
	err := authorUsecase.authorRepository.UpdateByAuthor(ctx, &authorEntity)
	if err != nil {
		return err
	}

	return nil
}

// 削除
func (authorUsecase *authorUsecase) Delete(ctx context.Context, author *domain.Author, userId int) error {

	var authorEntity entity.Author
	authorEntity.Id = author.GetId()
	authorEntity.Name = author.GetName()
	authorEntity.UserId = author.GetUserId()
	err := authorUsecase.authorRepository.DeleteByAuthor(ctx, &authorEntity, userId)
	if err != nil {
		return err
	}

	return nil
}

// カテゴリー取得チャネル
func (authorUsecase *authorUsecase) getAuthorByUserChannel(userId int, ch chan *channel.AuthorGetChannel) {

	var authors []domain.Author
	var authorEntities []entity.Author
	var err error
	authorGetChannel, err := channel.NewAuthorGetChannel(authors, err)

	authorByUsers, err := authorUsecase.authorRepository.GetAuthorByUser(authorEntities, userId)
	for _, author := range authorByUsers {
		var authorModel domain.Author
		authorModel.Set(author.Id, author.Name, author.UserId, author.UpdatedAt, author.CreatedAt)
		authors = append(authors, authorModel)
	}

	authorGetChannel.Set(authors, err)
	ch <- authorGetChannel
}

// 経験値・レベル取得チャネル
func (authorUsecase *authorUsecase) getExpByUserChannel(userId int, ch chan *channel.ExpPoolGetChannel) {

	var expPool domain.ExpPool
	var err error
	expPoolGetChannel, err := channel.NewExpPoolGetChannel(expPool, err)

	expPool, err = authorUsecase.expPoolUsecase.GetByUserId(userId)

	expPoolGetChannel.Set(expPool, err)
	ch <- expPoolGetChannel
}
