package account

import (
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/person"
)

// Entity - Account Entity.
type Entity struct {
	ID          uuid.UUID      `json:"id"`
	Email       Email          `json:"email"`
	Password    Password       `json:"password"`
	Active      bool           `json:"active"`
	LastLoginAt time.Time      `json:"lastLoginAt"`
	Person      *person.Entity `json:"person"`
}

// New creates a new account entity.
func New() Entity {
	p := person.New()
	return Entity{
		ID:     uuid.New(),
		Person: &p,
	}
}

// IsActive returns a boolean if the account is active or not.
func (e *Entity) IsActive() bool {
	return e.Active
}

// Activate sets the account entity active field to true.
func (e *Entity) Activate() {
	e.Active = true
}

// Deactivate sets the account entity active field to true.
func (e *Entity) Deactivate() {
	e.Active = false
}

// LoginTimestamp records the date time when the user logs in.
func (e *Entity) LoginTimestamp() {
	e.LastLoginAt = time.Now()
}

// UpdateEmail checks the validity of the email address
// and updates the account entity email field value.
func (e *Entity) UpdateEmail(s string) error {
	email, err := e.Email.Update(s)
	if err != nil {
		return err
	}

	e.Email = email
	return nil
}

// UpdatePassword checks the validity of the password
// and updates the account entity password field value.
func (e *Entity) UpdatePassword(s string) error {
	password, err := e.Password.Update(s)
	if err != nil {
		return err
	}

	e.Password = password
	return nil
}

// UpdatePersonAvatar takes a string parameter that should contain a valid avatar URL.
// If validation failed it will return an error,
// otherwise it will update the corresponding field in the person entity.
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
		idx := indexOf(input.ID, addrs)

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

func indexOf(id uuid.UUID, data []address.Entity) int {
	for k, v := range data {
		if id == v.ID {
			return k
		}
	}
	return -1
}
