package domain

import (
	"errors"
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

// constructor
func NewUser(id int, email, password string, updatedAt, createdAt time.Time) (*User, error) {
	user := &User{
		Id:        id,
		Email:     email,
		Password:  password,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}

	return user, nil
}

// setter
func (user *User) Set(id int, email, password string, updatedAt, createdAt time.Time) error {
	if email == "" {
		return errors.New("email is required")
	}
	user.Id = id
	user.Email = email
	user.Password = password
	user.UpdatedAt = updatedAt
	user.CreatedAt = createdAt

	return nil
}
