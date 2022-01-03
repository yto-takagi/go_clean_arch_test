package entity

import (
	"time"
)

type ExpPool struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Exp       int       `json:"exp"`
	Lv        int       `json:"lv"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
