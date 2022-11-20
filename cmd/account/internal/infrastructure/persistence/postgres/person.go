package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/person"
)

// Person postgres model.
type Person struct {
	ID          uuid.UUID `json:"id"`
	AccountID   uuid.UUID `json:"accountId"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Avatar      string    `json:"avatar"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ToEntity transforms the person model to account entity.
func (p Person) ToEntity() person.Entity {
	return person.Entity{
		ID:        p.ID,
		AccountID: p.AccountID,
		Details: person.Details{
			FirstName:   p.FirstName,
			LastName:    p.LastName,
			Email:       p.Email,
			Phone:       p.Phone,
			DateOfBirth: p.DateOfBirth,
		},
		Avatar: person.Avatar(p.Avatar),
	}
}
