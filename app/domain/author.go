package domain

import "time"

type Author struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	UserId    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
