package account

import (
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/person"
	"github.com/ssengalanto/potato-project/pkg/gg"
	"github.com/ssengalanto/potato-project/pkg/validator"
)

// Entity - account entity struct.
type Entity struct {
	ID          uuid.UUID      `json:"id" validate:"required"`
	Email       Email          `json:"email" validate:"required"`
	Password    Password       `json:"password" validate:"required"`
	Active      bool           `json:"active" validate:"required"`
	LastLoginAt time.Time      `json:"lastLoginAt,omitempty"`
	Person      *person.Entity `json:"person"`
}

// New creates a new account entity.
func New(email, password string, active bool) Entity {
	return Entity{
		ID:       uuid.New(),
		Email:    Email(email),
		Password: Password(password),
		Active:   active,
	}
}

// IsActive returns the active state of the account entity.
func (e *Entity) IsActive() bool {
	return e.Active
}

// Activate sets the active state of the account entity to `true`.
func (e *Entity) Activate() {
	e.Active = true
}

// Deactivate sets the active state of the account entity to `false`.
func (e *Entity) Deactivate() {
	e.Active = false
}

// LoginTimestamp sets the last login at state of the account entity with current timestamp.
func (e *Entity) LoginTimestamp() {
	e.LastLoginAt = time.Now()
}

// UpdateEmail checks the validity of the email address and updates the account entity email field value.
func (e *Entity) UpdateEmail(s string) error {
	email, err := e.Email.Update(s)
	if err != nil {
		return err
	}

	e.Email = email
	return nil
}

// UpdatePassword checks the validity of the password and updates the account entity password field value.
func (e *Entity) UpdatePassword(s string) error {
	password, err := e.Password.Update(s)
	if err != nil {
		return err
	}

	e.Password = password
	return nil
}

// UpdatePersonAvatar takes a string parameter that should contain a valid avatar URL.
// If validation failed it will return an error, otherwise it will update the corresponding field in the person entity.
func (e *Entity) UpdatePersonAvatar(avatar string) error {
	err := e.Person.UpdateAvatar(avatar)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePersonDetails takes a struct parameter that contains the person details
// to be used for the update. If validation failed it will return an error
// otherwise it will update the corresponding fields in person entity.
func (e *Entity) UpdatePersonDetails(input person.UpdateDetailsInput) error {
	err := e.Person.UpdateDetails(input)
	if err != nil {
		return err
	}

	return nil
}

// UpdateAddress takes a slice of struct parameter that contains the address components
// to be used for the update. If validation failed it will return an error
// otherwise it will update the corresponding fields in address entity.
func (e *Entity) UpdateAddress(inputs []address.UpdateInput) error {
	addrs := *e.Person.Address

	for _, input := range inputs {
		idx := gg.FindIndexOf[address.Entity](addrs, func(addr address.Entity) bool {
			return addr.ID == input.ID
		})

		if idx == -1 {
			continue
		}

		err := addrs[idx].Update(input)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Entity) IsValid() error {
	err := validator.Struct(e)
	if err != nil {
		return err
	}

	_, err = e.Email.IsValid()
	if err != nil {
		return err
	}

	_, err = e.Password.IsValid()
	if err != nil {
		return err
	}

	return err
}
