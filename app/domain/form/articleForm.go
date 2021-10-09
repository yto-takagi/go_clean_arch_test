package domain

import "time"

// struct
type ArticleForm struct {
	// gorm.ModelはID, CreatedAt, UpdatedAt, DeletedAtをフィールドに持つ構造体
	// gorm.Model
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	AuthorId int `json:"author_id"`
}
