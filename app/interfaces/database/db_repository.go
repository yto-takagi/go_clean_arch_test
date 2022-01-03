package database

import (
	"github.com/jinzhu/gorm"
)

type DBRepository struct {
	DB DB
}

func (db *DBRepository) Connect() *gorm.DB {
	return db.DB.Connect()
}
