package repository

import (
	"go_clean_arch_test/app/interfaces/database/repository/entity"
)

// LoginRepository interface
type LoginRepository interface {
	GetByEmail(email string, user entity.User) (entity.User, error)
}
