package domain

import (
	"errors"
	"time"
)

// struct
type User struct {
	id        int       `json:"id"`
	email     string    `json:"email"`
	password  string    `json:"password"`
	updatedAt time.Time `json:"updated_at"`
	createdAt time.Time `json:"created_at"`
}

// constructor
func NewUser(id int, email, password string, updatedAt, createdAt time.Time) (*User, error) {
	user := &User{
		id:        id,
		email:     email,
		password:  password,
		updatedAt: updatedAt,
		createdAt: createdAt,
	}

	return user, nil
}

// setter
func (user *User) Set(id int, email, password string, updatedAt, createdAt time.Time) error {
	if email == "" {
		return errors.New("email is required")
	}
	user.id = id
	user.email = email
	user.password = password
	user.updatedAt = updatedAt
	user.createdAt = createdAt

	return nil
}

// getter
func (user *User) GetId() int {
	return user.id
}

func (user *User) GetEmail() string {
	return user.email
}

func (user *User) GetPassword() string {
	return user.password
}

func (user *User) GetUpdatedAt() time.Time {
	return user.updatedAt
}

func (user *User) GetCreatedAt() time.Time {
	return user.createdAt
}
