package domain

import (
	"errors"
	"time"
)

// struct
type ExpPoolForm struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Exp       int       `json:"exp"`
	Lv        int       `json:"lv"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// constructor
func NewExpPoolForm(id, userId, exp, lv int, updatedAt, createdAt time.Time) (*ExpPoolForm, error) {
	expPoolForm := &ExpPoolForm{
		Id:        id,
		UserId:    userId,
		Exp:       exp,
		Lv:        lv,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}

	return expPoolForm, nil
}

// setter
func (expPoolForm *ExpPoolForm) Set(id, userId, exp, lv int, updatedAt, createdAt time.Time) error {
	if userId == 0 {
		return errors.New("userId is required")
	}
	expPoolForm.Id = id
	expPoolForm.UserId = userId
	expPoolForm.Exp = exp
	expPoolForm.Lv = lv
	expPoolForm.UpdatedAt = updatedAt
	expPoolForm.CreatedAt = createdAt

	return nil
}
