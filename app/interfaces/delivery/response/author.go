package response

import (
	"time"
)

type Author struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	UserId    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// construcor
func NewAuthor(id int, name string, userId int, updatedAt, createdAt time.Time) *Author {
	Author := new(Author)
	Author.Id = id
	Author.Name = name
	Author.UserId = userId
	Author.UpdatedAt = updatedAt
	Author.CreatedAt = createdAt
	return Author
}
