package response

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

// construcor
func NewExpPool(id, userId, exp, lv int, updatedAt, createdAt time.Time) *ExpPool {
	ExpPool := new(ExpPool)
	ExpPool.Id = id
	ExpPool.UserId = userId
	ExpPool.Exp = exp
	ExpPool.Lv = lv
	ExpPool.UpdatedAt = updatedAt
	ExpPool.CreatedAt = createdAt
	return ExpPool
}
