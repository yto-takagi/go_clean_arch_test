package sql

import (
	"errors"
	"go_clean_arch_test/app/interfaces/database/repository/entity"
	"go_clean_arch_test/app/usecase/repository"

	"github.com/jinzhu/gorm"
)

// LvRepository struct
type LvRepository struct {
	Conn *gorm.DB
}

// NewLvRepository constructor
func NewLvRepository(conn *gorm.DB) repository.LvRepository {
	return &LvRepository{Conn: conn}
}

// 必要経験値に紐付いたデータ取得
func (lvRepository *LvRepository) GetByExp(lv entity.Lv, exp int) (entity.Lv, error) {
	if err := lvRepository.Conn.
		Debug().
		Table("lv").
		Select("Max(lv.lv) AS lv").
		Where("lv.necessary <= ?", exp).
		Scan(&lv).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return lv, err
		}
	}

	return lv, nil

}
