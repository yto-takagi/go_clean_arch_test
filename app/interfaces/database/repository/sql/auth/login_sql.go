package auth

import (
	"go_clean_arch_test/app/interfaces/database/repository/entity"
	repository "go_clean_arch_test/app/usecase/repository/auth"
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
func (loginRepository *LoginRepository) GetByEmail(email string, user entity.User) (entity.User, error) {
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
	logger.Info("GetByEmail",
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	return user, nil
}
