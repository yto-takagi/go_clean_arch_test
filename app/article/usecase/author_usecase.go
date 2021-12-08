package usecase

import (
	"context"
	"go_clean_arch_test/app/domain"
	"go_clean_arch_test/app/domain/repository"
	"time"

	"go.uber.org/zap"
)

// AuthorUsecase interface
type AuthorUsecase interface {
	GetByUser(userId int) ([]domain.Author, error)
	GetByAuthorIdAndUserId(id int, userId int) (domain.Author, error)
	GetByName(name string, userId int) (domain.Author, error)
	Input(ctx context.Context, author *domain.Author) (domain.Author, error)
	Update(ctx context.Context, author *domain.Author) error
	Delete(ctx context.Context, author *domain.Author, userId int) error
}
type authorUsecase struct {
	authorRepository repository.AuthorRepository
}

// NewAuthorUsecase constructor
func NewAuthorUsecase(authorRepository repository.AuthorRepository) AuthorUsecase {
	return &authorUsecase{authorRepository: authorRepository}
}

// ユーザーIDに紐づくカテゴリー全件取得
func (authorUsecase *authorUsecase) GetByUser(userId int) ([]domain.Author, error) {

	var author []domain.Author
	authorByUser, err := authorUsecase.authorRepository.GetAuthorByUser(author, userId)
	if err != nil {
		return nil, err
	}

	return authorByUser, nil
}

// Id、ユーザーID指定
func (authorUsecase *authorUsecase) GetByAuthorIdAndUserId(id int, userId int) (domain.Author, error) {

	var author domain.Author
	authorByIdAndUserId, err := authorUsecase.authorRepository.GetAuthorByAuthorIdAndUserId(author, id, userId)
	if err != nil {
		return author, err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("GetById",
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return authorByIdAndUserId, nil
}

// カテゴリー名検索
func (authorUsecase *authorUsecase) GetByName(name string, userId int) (domain.Author, error) {

	var author domain.Author
	authorByName, err := authorUsecase.authorRepository.GetByAuthorName(author, name, userId)
	if err != nil {
		return author, err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("GetByName",
		zap.String("param name", name),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return authorByName, nil
}

// 新規登録
func (authorUsecase *authorUsecase) Input(ctx context.Context, author *domain.Author) (domain.Author, error) {

	author.Id = 0
	author.CreatedAt = time.Now()
	author.UpdatedAt = time.Now()

	err := authorUsecase.authorRepository.InputByAuthor(ctx, author)
	if err != nil {
		return *author, err
	}

	return *author, nil
}

// 更新
func (authorUsecase *authorUsecase) Update(ctx context.Context, author *domain.Author) error {

	author.UpdatedAt = time.Now()
	err := authorUsecase.authorRepository.UpdateByAuthor(ctx, author)
	if err != nil {
		return err
	}

	return nil
}

// 削除
func (authorUsecase *authorUsecase) Delete(ctx context.Context, author *domain.Author, userId int) error {

	err := authorUsecase.authorRepository.DeleteByAuthor(ctx, author, userId)
	if err != nil {
		return err
	}

	return nil
}
