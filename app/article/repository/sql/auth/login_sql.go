package auth

import (
	domain "go_clean_arch_test/app/domain"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// emailに紐付いたデータ取得
func GetByEmail(db *gorm.DB, email string, user domain.User) domain.User {
	db.
		Debug().
		Table("users").
		Select("users.id, users.email, users.password, users.created_at, users.updated_at").
		Where("users.email = ?", email).
		Scan(&user)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ login_sql.go ++++++++++++++++++++++",
		zap.String("method", "GetByEmail"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(user)

	return user
}
