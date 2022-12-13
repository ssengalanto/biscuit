package person

import (
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
	"github.com/ssengalanto/potato-project/pkg/validator"
)

// Entity - Person Entity.
type Entity struct {
	ID        uuid.UUID         `json:"id" validate:"uuid,required"`
	AccountID uuid.UUID         `json:"accountId" validate:"uuid,required"`
	Details   Details           `json:"details" validate:"required"`
	Avatar    Avatar            `json:"avatar"`
	Address   *[]address.Entity `json:"address"`
}

// UpdateDetailsInput - input for updating peron Details.
type UpdateDetailsInput struct {
	FirstName   *string
	LastName    *string
	Email       *string
	Phone       *string
	DateOfBirth *time.Time
}

// New creates a new person entity.
func New(accountID uuid.UUID, firstName, lastName, email, phone string, dateOfBirth time.Time) Entity {
	return Entity{
		ID:        uuid.New(),
		AccountID: accountID,
		Details: Details{
			FirstName:   firstName,
			LastName:    lastName,
			Email:       email,
			Phone:       phone,
			DateOfBirth: dateOfBirth,
		},
	}
}

func (e *Entity) UpdateDetails(input UpdateDetailsInput) error {
	details := e.Details

	if input.FirstName != nil {
		details.FirstName = *input.FirstName
	}

	if input.LastName != nil {
		details.LastName = *input.LastName
	}

	if input.Email != nil {
		details.Email = *input.Email
	}

	if input.Phone != nil {
		details.Phone = *input.Phone
	}

	if input.DateOfBirth != nil {
		details.DateOfBirth = *input.DateOfBirth
	}

	newDetails, err := e.Details.Update(details)
	if err != nil {
		return err
	}

	e.Details = newDetails
	return nil
}

func (e *Entity) UpdateAvatar(s string) error {
	avatar, err := e.Avatar.Update(s)
	if err != nil {
		return err
	}

	e.Avatar = avatar
	return nil
}

func (e *Entity) IsValid() error {
	err := validator.Struct(e)
	if err != nil {
		return err
	}

	err = e.Details.IsValid()
	if err != nil {
		return err
	}

	return err
}
