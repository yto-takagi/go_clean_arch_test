package auth

import (
	"errors"
	"go_clean_arch_test/app/domain"
	form "go_clean_arch_test/app/domain/form"
	repository "go_clean_arch_test/app/domain/repository/auth"
	"log"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// SignUpUsecase interface
type SignUpUsecase interface {
	SignUp(email string, password string) (domain.User, error)
}

type signUpUsecase struct {
	signUpRepository  repository.SignUpRepository
	loginUpRepository repository.LoginRepository
}

// NewSignUpUsecase constructor
func NewSignUpUsecase(signUpRepository repository.SignUpRepository, loginUpRepository repository.LoginRepository) SignUpUsecase {
	return &signUpUsecase{signUpRepository: signUpRepository, loginUpRepository: loginUpRepository}
}

var (
	ErrEmailIsExists = errors.New("email is exists")
	ErrUnknown       = errors.New("unknown error")
)

// 会員登録
func (signUpUsecase *signUpUsecase) SignUp(email string, password string) (domain.User, error) {

	var user domain.User
	var signUp form.SignUpForm
	signUp.Email = email
	signUp.Password = password

	// email存在チェック
	user, err := signUpUsecase.loginUpRepository.GetByEmail(email, user)
	if err != nil {
		return user, err
	}
	if user.Id != 0 {
		return user, ErrEmailIsExists
	}

	// 会員登録
	// パスワードハッシュ化
	hashed, _ := bcrypt.GenerateFromPassword([]byte(signUp.Password), 10)
	signUp.Password = string(hashed)

	var userInfo domain.User
	err = signUpUsecase.signUpRepository.SignUp(&signUp)
	if err != nil {
		return userInfo, err
	}

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
