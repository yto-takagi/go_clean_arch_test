package request

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
