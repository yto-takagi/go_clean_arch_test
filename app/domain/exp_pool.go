package domain

import (
	"errors"
	"time"
)

type ExpPool struct {
	id        int       `json:"id"`
	userId    int       `json:"user_id"`
	exp       int       `json:"exp"`
	lv        int       `json:"lv"`
	updatedAt time.Time `json:"updated_at"`
	createdAt time.Time `json:"created_at"`
}

// constructor
func NewExpPool(id, userId, exp, lv int, updatedAt, createdAt time.Time) (*ExpPool, error) {
	expPool := &ExpPool{
		id:        id,
		userId:    userId,
		exp:       exp,
		lv:        lv,
		updatedAt: updatedAt,
		createdAt: createdAt,
	}

	return expPool, nil
}

// setter
func (expPool *ExpPool) Set(id, userId, exp, lv int, updatedAt, createdAt time.Time) error {
	if userId == 0 {
		return errors.New("userId is required")
	}
	expPool.id = id
	expPool.userId = userId
	expPool.exp = exp
	expPool.lv = lv
	expPool.updatedAt = updatedAt
	expPool.createdAt = createdAt

	return nil
}

// getter
func (expPool *ExpPool) GetId() int {
	return expPool.id
}

func (expPool *ExpPool) GetUserId() int {
	return expPool.userId
}

func (expPool *ExpPool) GetExp() int {
	return expPool.exp
}

func (expPool *ExpPool) GetLv() int {
	return expPool.lv
}

func (expPool *ExpPool) GetUpdatedAt() time.Time {
	return expPool.updatedAt
}

func (expPool *ExpPool) GetCreatedAt() time.Time {
	return expPool.createdAt
}
