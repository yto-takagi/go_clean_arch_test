package auth

import (
	"errors"
	sql "go_clean_arch_test/app/article/repository/sql/auth"
	"go_clean_arch_test/app/article/usecase"
	"go_clean_arch_test/app/domain"
	form "go_clean_arch_test/app/domain/form"
	"log"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type SignUpUsecase struct {
	DB usecase.DBRepository
}

var (
	ErrEmailIsExists = errors.New("email is exists")
	ErrUnknown       = errors.New("unknown error")
)

// 会員登録
func (usecase *SignUpUsecase) SignUp(email string, password string) (userInfo domain.User, err error) {
	db := usecase.DB.Connect()
	// defer db.Close()

	var user domain.User
	var signUp form.SignUpForm
	signUp.Email = email
	signUp.Password = password

	// email存在チェック
	user = sql.GetByEmail(db, email, user)
	if user.Id != 0 {
		return user, ErrEmailIsExists
	}

	// 会員登録
	// パスワードハッシュ化対応 新規登録時使用
	hashed, _ := bcrypt.GenerateFromPassword([]byte(signUp.Password), 10)
	signUp.Password = string(hashed)

	sql.SignUp(db, &signUp)
	userInfo.Id = signUp.Id
	userInfo.Email = signUp.Email
	userInfo.Password = signUp.Password
	userInfo.CreatedAt = signUp.CreatedAt
	userInfo.UpdatedAt = signUp.UpdatedAt

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ signup_usecase.go ++++++++++++++++++++++",
		zap.String("method", "SignUp"),
		zap.String("param email", email),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■会員登録 userInfo■■■■■■■■■■■■■■■■■■■■■■■■■■")
	log.Println(userInfo)

	return userInfo, nil
}
