package repository

import (
	"go_clean_arch_test/app/domain"
)

// LoginRepository interface
type LoginRepository interface {
	GetByEmail(email string, user domain.User) (domain.User, error)
}
