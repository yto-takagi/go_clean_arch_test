package auth

import (
	"errors"
	"time"
)

// struct
type Login struct {
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// constructor
func NewLogin(email, password string, updatedAt, createdAt time.Time) (*Login, error) {

	login := &Login{
		Email:     email,
		Password:  password,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}

	return login, nil
}

// setter
func (login *Login) Set(email, password string, updatedAt, createdAt time.Time) error {

	if email == "" || password == "" {
		return errors.New("email and password is required")
	}
	login.Email = email
	login.Password = password
	login.UpdatedAt = updatedAt
	login.CreatedAt = createdAt

	return nil
}
