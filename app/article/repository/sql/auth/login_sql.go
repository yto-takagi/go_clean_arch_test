package auth

import (
	domain "go_clean_arch_test/app/domain"
	repository "go_clean_arch_test/app/domain/repository/auth"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type LoginRepository struct {
	Conn *gorm.DB
}

// NewLoginRepository constructor
func NewLoginRepository(conn *gorm.DB) repository.LoginRepository {
	return &LoginRepository{Conn: conn}
}

// emailに紐付いたデータ取得
func (loginRepository *LoginRepository) GetByEmail(email string, user domain.User) (domain.User, error) {
	if err := loginRepository.Conn.
		Debug().
		Table("users").
		Select("users.id, users.email, users.password, users.created_at, users.updated_at").
		Where("users.email = ?", email).
		Scan(&user).
		Error; err != nil {
		return user, err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ login_sql.go ++++++++++++++++++++++",
		zap.String("method", "GetByEmail"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(user)

	return user, nil
}
