package auth

import (
	form "go_clean_arch_test/app/domain/form"
	repository "go_clean_arch_test/app/usecase/repository/auth"

	"github.com/jinzhu/gorm"
)

type SignUpRepository struct {
	Conn *gorm.DB
}

// NewSignUpRepository constructor
func NewSignUpRepository(conn *gorm.DB) repository.SignUpRepository {
	return &SignUpRepository{Conn: conn}
}

// 会員登録
func (signUpRepository *SignUpRepository) SignUp(signUpForm *form.SignUpForm) error {
	if err := signUpRepository.Conn.
		Debug().
		Table("users").
		Create(&signUpForm).
		Error; err != nil {
		return err
	}

	return nil
}
