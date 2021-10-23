package domain

import "time"

// struct
type User struct {
	// gorm.ModelはID, CreatedAt, UpdatedAt, DeletedAtをフィールドに持つ構造体
	// gorm.Model
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
