package usecase

import (
	"go_clean_arch_test/app/article/repository/sql"
	"go_clean_arch_test/app/domain"
	"log"
	"time"

	"go.uber.org/zap"
)

type AuthorUsecase struct {
	DB DBRepository
}

// ユーザーIDに紐づくカテゴリー全件取得
func (usecase *AuthorUsecase) GetByUser(userId int) []domain.Author {
	db := usecase.DB.Connect()
	// defer db.Close()

	var author []domain.Author
	authorByUser := sql.GetAuthorByUser(db, author, userId)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetByUser"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(authorByUser)

	return authorByUser
}

// Id、ユーザーID指定
func (usecase *AuthorUsecase) GetByAuthorIdAndUserId(id int, userId int) domain.Author {
	db := usecase.DB.Connect()
	// defer db.Close()

	var author domain.Author
	authorByIdAndUserId := sql.GetAuthorByAuthorIdAndUserId(db, author, id, userId)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetById"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(authorByIdAndUserId)

	return authorByIdAndUserId
}

// カテゴリー名検索
func (usecase *AuthorUsecase) GetByName(name string, userId int) domain.Author {
	db := usecase.DB.Connect()
	// defer db.Close()

	var author domain.Author
	authorByName := sql.GetByAuthorName(db, author, name, userId)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetByName"),
		zap.String("param name", name),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(authorByName)

	return authorByName
}

// 新規登録
func (usecase *AuthorUsecase) Input(author *domain.Author) domain.Author {
	db := usecase.DB.Connect()
	// defer db.Close()

	author.Id = 0
	author.CreatedAt = time.Now()
	author.UpdatedAt = time.Now()

	sql.InputByAuthor(db, author)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_usecase.go ++++++++++++++++++++++",
		zap.String("method", "Input"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(author)
	return *author
}

// 更新
func (usecase *AuthorUsecase) Update(author *domain.Author) {
	db := usecase.DB.Connect()
	// defer db.Close()

	author.UpdatedAt = time.Now()

	sql.UpdateByAuthor(db, author)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_usecase.go ++++++++++++++++++++++",
		zap.String("method", "Input"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(author)
}

// 削除
func (usecase *AuthorUsecase) Delete(author *domain.Author, userId int) {
	db := usecase.DB.Connect()
	// defer db.Close()

	sql.DeleteByAuthor(db, author, userId)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_usecase.go ++++++++++++++++++++++",
		zap.String("method", "Delete"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
}
