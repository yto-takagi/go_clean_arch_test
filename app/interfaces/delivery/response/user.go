package response

import (
	"time"
)

// struct
type User struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// construcor
func NewUser(id int, email, password string, updatedAt, createdAt time.Time) *User {
	User := new(User)
	User.Id = id
	User.Email = email
	User.Password = password
	User.UpdatedAt = updatedAt
	User.CreatedAt = createdAt
	return User
}
