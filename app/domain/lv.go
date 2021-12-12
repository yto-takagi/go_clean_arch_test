package domain

import (
	"time"
)

type Lv struct {
	id        int       `json:"id"`
	lv        int       `json:"lv"`
	necessary int       `json:"necessary"`
	updatedAt time.Time `json:"updated_at"`
	createdAt time.Time `json:"created_at"`
}

// constructor
func NewLv(id, lv, necessary int, updatedAt, createdAt time.Time) (*Lv, error) {
	lvModel := &Lv{
		id:        id,
		lv:        lv,
		necessary: necessary,
		updatedAt: updatedAt,
		createdAt: createdAt,
	}

	return lvModel, nil
}

// setter
func (lvModel *Lv) Set(id, lv, necessary int, updatedAt, createdAt time.Time) error {
	lvModel.id = id
	lvModel.lv = lv
	lvModel.necessary = necessary
	lvModel.updatedAt = updatedAt
	lvModel.createdAt = createdAt

	return nil
}

// getter
func (lvModel *Lv) GetId() int {
	return lvModel.id
}

func (lvModel *Lv) GetLv() int {
	return lvModel.lv
}

func (lvModel *Lv) GetNecessary() int {
	return lvModel.necessary
}

func (lvModel *Lv) GetUpdatedAt() time.Time {
	return lvModel.updatedAt
}

func (lvModel *Lv) GetCreatedAt() time.Time {
	return lvModel.createdAt
}
