package auth

import "errors"

// struct
type SignUp struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

// constructor
func NewSignUp(email, password string) (*SignUp, error) {

	signUp := &SignUp{
		Email:    email,
		Password: password,
	}

	return signUp, nil
}

// setter
func (signUp *SignUp) Set(email, password string) error {

	if email == "" || password == "" {
		return errors.New("email and password is required")
	}
	signUp.Email = email
	signUp.Password = password

	return nil
}
