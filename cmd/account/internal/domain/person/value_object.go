package person

import (
	"time"

	"github.com/ssengalanto/potato-project/pkg/validator"
)

// Details - person details value object.
type Details struct {
	FirstName   string    `json:"firstName" validate:"required"`
	LastName    string    `json:"lastName" validate:"required"`
	Email       string    `json:"email" validate:"email,required"`
	Phone       string    `json:"phone" validate:"required"`
	DateOfBirth time.Time `json:"dateOfBirth" validate:"required"`
}

// IsValid checks the validity of the person details.
func (d Details) IsValid() error {
	err := validator.Struct(d)
	if err != nil {
		return err
	}

	return err
}

func (d Details) Update(input Details) (Details, error) {
	err := input.IsValid()
	if err != nil {
		return Details{}, err
	}

	return input, nil
}

// Avatar value object.
type Avatar string

// IsValid checks the validity of the avatar url.
func (a Avatar) IsValid() (bool, error) {
	err := validator.Var("Avatar", a, "url,required")
	if err != nil {
		return false, err
	}

	return true, nil
}

// Update checks the validity of the avatar url and updates its value.
func (a Avatar) Update(s string) (Avatar, error) {
	avatar := Avatar(s)
	if ok, err := avatar.IsValid(); !ok {
		return "", err
	}

	return avatar, nil
}

// String converts Avatar to type string.
func (a Avatar) String() string {
	return string(a)
}
