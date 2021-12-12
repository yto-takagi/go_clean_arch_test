package response

import (
	"time"
)

type Lv struct {
	Id        int       `json:"id"`
	Lv        int       `json:"lv"`
	Necessary int       `json:"necessary"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// construcor
func NewLv(id, lv, necessary int, updatedAt, createdAt time.Time) *Lv {
	Lv := new(Lv)
	Lv.Id = id
	Lv.Lv = lv
	Lv.Necessary = necessary
	Lv.UpdatedAt = updatedAt
	Lv.CreatedAt = createdAt
	return Lv
}
