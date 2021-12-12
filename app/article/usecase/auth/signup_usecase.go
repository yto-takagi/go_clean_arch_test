package auth

import (
	"errors"
	"go_clean_arch_test/app/article/repository/entity"
	"go_clean_arch_test/app/domain"
	form "go_clean_arch_test/app/domain/form"
	repository "go_clean_arch_test/app/domain/repository/auth"
	"time"

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

	var user entity.User
	var userInfo domain.User
	signUp, err := form.NewSignUpForm(0, email, password, time.Now(), time.Now())
	if err != nil {
		return userInfo, err
	}

	// email存在チェック
	user, _ = signUpUsecase.loginUpRepository.GetByEmail(email, user)
	if userInfo.GetId() != 0 {
		return userInfo, ErrEmailIsExists
	}

	// 会員登録
	// パスワードハッシュ化
	hashed, _ := bcrypt.GenerateFromPassword([]byte(signUp.Password), 10)
	signUp.Set(signUp.Id, signUp.Email, string(hashed), signUp.UpdatedAt, signUp.CreatedAt)

	err = signUpUsecase.signUpRepository.SignUp(signUp)
	if err != nil {
		return userInfo, err
	}

	userInfo.Set(signUp.Id, signUp.Email, signUp.Password, signUp.UpdatedAt, signUp.CreatedAt)

	return userInfo, nil
}
