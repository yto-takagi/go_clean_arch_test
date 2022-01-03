package usecase

import (
	"go_clean_arch_test/app/domain"
	"go_clean_arch_test/app/interfaces/database/repository/entity"
	"go_clean_arch_test/app/transaction"
	"go_clean_arch_test/app/usecase/repository"
)

// LvUsecase interface
type LvUsecase interface {
	GetByExp(exp int) (domain.Lv, error)
}

type lvUsecase struct {
	lvRepository repository.LvRepository
	trancaction  transaction.Transaction
}

// NewLvUsecase constructor
func NewLvUsecase(lvRepository repository.LvRepository, trancaction transaction.Transaction) LvUsecase {
	return &lvUsecase{lvRepository: lvRepository, trancaction: trancaction}
}

// 累計経験値を元にデータ取得
func (lvUsecase *lvUsecase) GetByExp(exp int) (domain.Lv, error) {

	var lv entity.Lv
	var lvModel domain.Lv
	lv, err := lvUsecase.lvRepository.GetByExp(lv, exp)
	if err != nil {
		return lvModel, err
	}
	lvModel.Set(lv.Id, lv.Lv, lv.Necessary, lv.UpdatedAt, lv.CreatedAt)

	return lvModel, nil
}
