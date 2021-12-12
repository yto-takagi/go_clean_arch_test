package repository

import (
	"context"
	"go_clean_arch_test/app/article/repository/entity"
	form "go_clean_arch_test/app/domain/form"
)

// ExpPoolRepository interface
type ExpPoolRepository interface {
	GetByUserId(expPool entity.ExpPool, userId int) (entity.ExpPool, error)
	Input(ctx context.Context, expPoolForm *form.ExpPoolForm) error
	Update(ctx context.Context, expPoolForm *form.ExpPoolForm) error
}
