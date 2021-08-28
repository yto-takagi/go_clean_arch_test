package domain

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	Name string `json:"name"`
}
