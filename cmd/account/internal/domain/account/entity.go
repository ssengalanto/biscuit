package account

import (
	"time"

	"github.com/google/uuid"
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
	return Entity{
		ID: uuid.New(),
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
func (e *Entity) UpdateEmail(email string) error {
	newEmail, err := e.Email.Update(email)
	if err != nil {
		return err
	}

	e.Email = newEmail
	return nil
}

// UpdatePassword checks the validity of the password
// and updates the account entity password field value.
func (e *Entity) UpdatePassword(password string) error {
	newPassword, err := e.Password.Update(password)
	if err != nil {
		return err
	}

	e.Password = newPassword
	return nil
}
