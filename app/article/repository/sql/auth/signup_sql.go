package auth

import (
	form "go_clean_arch_test/app/domain/form"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// 会員登録
func SignUp(db *gorm.DB, signUpForm *form.SignUpForm) {
	db.
		Debug().
		Table("users").
		Create(&signUpForm)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ signup_sql.go ++++++++++++++++++++++",
		zap.String("method", "SignUp"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(signUpForm)

}
