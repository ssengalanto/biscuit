package person

import (
	"time"

	"github.com/ssengalanto/potato-project/pkg/validator"
)

// Details - person details value object.
type Details struct {
	FirstName   string    `validate:"required"`
	LastName    string    `validate:"required"`
	Email       string    `validate:"email,required"`
	Phone       string    `validate:"required"`
	DateOfBirth time.Time `validate:"required"`
}

// IsValid checks the validity of the person details.
func (d Details) IsValid() (bool, error) {
	err := validator.Struct(d)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (d Details) Update(input Details) (Details, error) {
	_, err := input.IsValid()
	if err != nil {
		return Details{}, err
	}

	return input, nil
}