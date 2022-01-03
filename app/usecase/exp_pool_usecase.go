package usecase

import (
	"context"
	"go_clean_arch_test/app/domain"
	form "go_clean_arch_test/app/domain/form"
	"go_clean_arch_test/app/interfaces/database/repository/entity"
	"go_clean_arch_test/app/transaction"
	"go_clean_arch_test/app/usecase/repository"
	"time"
)

// ExpPoolUsecase interface
type ExpPoolUsecase interface {
	GetByUserId(userId int) (domain.ExpPool, error)
	Input(ctx context.Context, expPool *domain.ExpPool) error
	Update(ctx context.Context, expPool *domain.ExpPool) error
}

type expPoolUsecase struct {
	expPoolRepository repository.ExpPoolRepository
	trancaction       transaction.Transaction
}

// NewExpPoolUsecase constructor
func NewExpPoolUsecase(expPoolRepository repository.ExpPoolRepository, trancaction transaction.Transaction) ExpPoolUsecase {
	return &expPoolUsecase{expPoolRepository: expPoolRepository, trancaction: trancaction}
}

// ユーザーID指定
func (expPoolUsecase *expPoolUsecase) GetByUserId(userId int) (domain.ExpPool, error) {

	var expPool entity.ExpPool
	var expPoolModel domain.ExpPool
	expPool, err := expPoolUsecase.expPoolRepository.GetByUserId(expPool, userId)
	if err != nil {
		return expPoolModel, err
	}
	expPoolModel.Set(expPool.Id, expPool.UserId, expPool.Exp, expPool.Lv, expPool.UpdatedAt, expPool.CreatedAt)

	return expPoolModel, nil
}

// 新規登録
// トランザクション
func (expPoolUsecase *expPoolUsecase) Input(ctx context.Context, expPool *domain.ExpPool) error {

	_, err := expPoolUsecase.trancaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		expPoolForm, err := form.NewExpPoolForm(0, expPool.GetUserId(), expPool.GetExp(), expPool.GetLv(), time.Now(), time.Now())
		if err != nil {
			return expPool, err
		}

		err = expPoolUsecase.expPoolRepository.Input(ctx, expPoolForm)
		if err != nil {
			return expPool, err
		}
		return expPool, nil
	})

	return err
}

// 更新
// トランザクション
func (expPoolUsecase *expPoolUsecase) Update(ctx context.Context, expPool *domain.ExpPool) error {

	_, err := expPoolUsecase.trancaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		expPoolForm, err := form.NewExpPoolForm(0, expPool.GetUserId(), expPool.GetExp(), expPool.GetLv(), expPool.GetUpdatedAt(), expPool.GetCreatedAt())
		if err != nil {
			return expPool, err
		}

		err = expPoolUsecase.expPoolRepository.Update(ctx, expPoolForm)
		if err != nil {
			return expPool, err
		}
		return expPool, nil
	})

	return err
}
