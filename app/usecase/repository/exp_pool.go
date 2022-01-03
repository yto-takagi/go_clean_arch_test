package repository

import (
	"context"
	form "go_clean_arch_test/app/domain/form"
	"go_clean_arch_test/app/interfaces/database/repository/entity"
)

// ExpPoolRepository interface
type ExpPoolRepository interface {
	GetByUserId(expPool entity.ExpPool, userId int) (entity.ExpPool, error)
	Input(ctx context.Context, expPoolForm *form.ExpPoolForm) error
	Update(ctx context.Context, expPoolForm *form.ExpPoolForm) error
}
