package domain

import (
	"errors"
	"time"
)

// struct
type SignUpForm struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// constructor
func NewSignUpForm(id int, email, password string, updatedAt, createdAt time.Time) (*SignUpForm, error) {

	signUpForm := &SignUpForm{
		Id:        id,
		Email:     email,
		Password:  password,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}

	return signUpForm, nil
}

// setter
func (signUpForm *SignUpForm) Set(id int, email, password string, updatedAt, createdAt time.Time) error {

	if email == "" || password == "" {
		return errors.New("email and password is required")
	}
	signUpForm.Id = id
	signUpForm.Email = email
	signUpForm.Password = password
	signUpForm.UpdatedAt = updatedAt
	signUpForm.CreatedAt = createdAt

	return nil
}
