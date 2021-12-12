package repository

import "go_clean_arch_test/app/article/repository/entity"

// LvRepository interface
type LvRepository interface {
	GetByExp(lv entity.Lv, exp int) (entity.Lv, error)
}
