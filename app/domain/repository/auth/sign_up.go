package repository

import (
	form "go_clean_arch_test/app/domain/form"
)

// SignUpRepository interface
type SignUpRepository interface {
	SignUp(signUpForm *form.SignUpForm) error
}
