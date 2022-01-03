package usecase

import (
	"github.com/jinzhu/gorm"
)

type DBRepository interface {
	Connect() *gorm.DB
}
