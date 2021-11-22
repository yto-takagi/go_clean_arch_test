package auth

import (
	form "go_clean_arch_test/app/domain/form"
	repository "go_clean_arch_test/app/domain/repository/auth"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
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

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ signup_sql.go ++++++++++++++++++++++",
		zap.String("method", "SignUp"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(signUpForm)

	return nil
}
