package auth

import (
	sql "go_clean_arch_test/app/article/repository/sql/auth"
	"go_clean_arch_test/app/article/usecase"
	"go_clean_arch_test/app/domain"
	auth "go_clean_arch_test/app/domain/auth"
	"log"
	"time"

	"go.uber.org/zap"
)

type LoginUsecase struct {
	DB usecase.DBRepository
}

// カテゴリー名検索
func (usecase *LoginUsecase) GetByEmail(email string) domain.User {
	db := usecase.DB.Connect()
	// defer db.Close()

	var user domain.User
	var login auth.Login
	login.Email = email
	userInfo := sql.GetByEmail(db, login, user)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetByName"),
		zap.String("param email", email),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(userInfo)

	return userInfo
}
