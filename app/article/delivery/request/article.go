package request

import (
	"time"
)

// struct
type Article struct {
	// gorm.Model
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	Author    `json:"author"`
	ExpPool   `json:"exp_pool"`
}
