package account

import (
	"github.com/ssengalanto/potato-project/pkg/validator"
)

type Email string

// IsValid checks the validity of the email address.
func (e Email) IsValid() (bool, error) {
	err := validator.Var("Email", e, "email,required")
	if err != nil {
		return false, err
	}

	return true, nil
}

// Update checks the validity of the email and updates its value.
func (e Email) Update(email string) (Email, error) {
	newEmail := Email(email)
	if ok, err := newEmail.IsValid(); !ok {
		return "", err
	}

	return newEmail, nil
}

// String converts type Email to type string.
func (e Email) String() string {
	return string(e)
}

type Password string

// IsValid checks the validity of the password.
func (p Password) IsValid() (bool, error) {
	err := validator.Var("Password", p, "min=10,required")
	if err != nil {
		return false, err
	}

	return true, nil
}

// Update checks the validity of the password and updates its value.
func (p Password) Update(password string) (Password, error) {
	newPassword := Password(password)
	if ok, err := newPassword.IsValid(); !ok {
		return "", err
	}

	return newPassword, nil
}

// String converts type Password to type string.
func (p Password) String() string {
	return string(p)
}
