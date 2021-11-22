package usecase

import (
	"go_clean_arch_test/app/domain"
	"go_clean_arch_test/app/domain/repository"
	"log"
	"time"

	"go.uber.org/zap"
)

// AuthorUsecase interface
type AuthorUsecase interface {
	GetByUser(userId int) ([]domain.Author, error)
	GetByAuthorIdAndUserId(id int, userId int) (domain.Author, error)
	GetByName(name string, userId int) (domain.Author, error)
	Input(author *domain.Author) (domain.Author, error)
	Update(author *domain.Author) error
	Delete(author *domain.Author, userId int) error
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

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetByUser"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(authorByUser)

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
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetById"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(authorByIdAndUserId)

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
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetByName"),
		zap.String("param name", name),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(authorByName)

	return authorByName, nil
}

// 新規登録
func (authorUsecase *authorUsecase) Input(author *domain.Author) (domain.Author, error) {

	author.Id = 0
	author.CreatedAt = time.Now()
	author.UpdatedAt = time.Now()

	err := authorUsecase.authorRepository.InputByAuthor(author)
	if err != nil {
		return *author, err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_usecase.go ++++++++++++++++++++++",
		zap.String("method", "Input"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(author)
	return *author, nil
}

// 更新
func (authorUsecase *authorUsecase) Update(author *domain.Author) error {

	author.UpdatedAt = time.Now()
	err := authorUsecase.authorRepository.UpdateByAuthor(author)
	if err != nil {
		return err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_usecase.go ++++++++++++++++++++++",
		zap.String("method", "Input"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(author)

	return nil
}

// 削除
func (authorUsecase *authorUsecase) Delete(author *domain.Author, userId int) error {

	err := authorUsecase.authorRepository.DeleteByAuthor(author, userId)
	if err != nil {
		return err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_usecase.go ++++++++++++++++++++++",
		zap.String("method", "Delete"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return nil
}
