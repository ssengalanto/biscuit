package account

import (
	"github.com/ssengalanto/potato-project/pkg/validator"
)

// Email value object.
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
func (e Email) Update(s string) (Email, error) {
	email := Email(s)
	if ok, err := email.IsValid(); !ok {
		return "", err
	}

	return email, nil
}

// String converts Email to type string.
func (e Email) String() string {
	return string(e)
}

// Password value object.
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
func (p Password) Update(s string) (Password, error) {
	password := Password(s)
	if ok, err := password.IsValid(); !ok {
		return "", err
	}

	return password, nil
}

// String converts Password to type string.
func (p Password) String() string {
	return string(p)
}
