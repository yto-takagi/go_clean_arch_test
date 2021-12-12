package sql

import (
	"context"
	"errors"
	"go_clean_arch_test/app/article/database"
	"go_clean_arch_test/app/article/repository/entity"
	form "go_clean_arch_test/app/domain/form"
	"go_clean_arch_test/app/domain/repository"

	"github.com/jinzhu/gorm"
)

// ExpPoolRepository struct
type ExpPoolRepository struct {
	Conn *gorm.DB
}

// NewExpPoolRepository constructor
func NewExpPoolRepository(conn *gorm.DB) repository.ExpPoolRepository {
	return &ExpPoolRepository{Conn: conn}
}

// ユーザーIdに紐付いたデータ取得
func (expPoolRepository *ExpPoolRepository) GetByUserId(expPool entity.ExpPool, userId int) (entity.ExpPool, error) {
	if err := expPoolRepository.Conn.
		Debug().
		Table("exp_pool").
		Select("exp_pool.id , exp_pool.user_id, exp_pool.exp, exp_pool.lv, exp_pool.created_at, exp_pool.updated_at").
		Where("exp_pool.user_id = ?", userId).
		Scan(&expPool).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return expPool, err
		}
	}

	return expPool, nil

}

// 新規登録
func (expPoolRepository *ExpPoolRepository) Input(ctx context.Context, expPoolForm *form.ExpPoolForm) error {
	dao, ok := database.GetTx(ctx)
	if !ok {
		dao = expPoolRepository.Conn
	}
	if err := dao.
		Debug().
		Table("exp_pool").
		Create(&expPoolForm).
		Error; err != nil {
		return err
	}

	return nil

}

// 更新
func (expPoolRepository *ExpPoolRepository) Update(ctx context.Context, expPoolForm *form.ExpPoolForm) error {
	dao, ok := database.GetTx(ctx)
	if !ok {
		dao = expPoolRepository.Conn
	}
	if err := dao.
		Debug().
		Model(&expPoolForm).
		Table("exp_pool").
		Omit("createdAt").
		Where("user_id = ?", expPoolForm.UserId).
		Updates(map[string]interface{}{
			"exp":        expPoolForm.Exp,
			"lv":         expPoolForm.Lv,
			"updated_at": expPoolForm.UpdatedAt}).
		Error; err != nil {
		return err
	}

	return nil
}
