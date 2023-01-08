package address

import (
	"github.com/google/uuid"
	"github.com/ssengalanto/hex/pkg/validator"
)

// Entity - address entity struct.
type Entity struct {
	ID         uuid.UUID  `json:"id"`
	PersonID   uuid.UUID  `json:"personId"`
	Components Components `json:"components"`
}

// New creates a new address entity.
func New(personID uuid.UUID, components Components) Entity {
	return Entity{
		ID:         uuid.New(),
		PersonID:   personID,
		Components: components,
	}
}

// UpdateComponentsInput contains required fields for updating address components.
type UpdateComponentsInput struct {
	Street     *string
	Unit       *string
	City       *string
	District   *string
	State      *string
	Country    *string
	PostalCode *string
}

// UpdateComponents checks the validity of the update components input
// and updates the address entity components fields.
func (e *Entity) UpdateComponents(input UpdateComponentsInput) error {
	components := e.Components

	if input.Street != nil {
		components.Street = *input.Street
	}

	if input.Unit != nil {
		components.Unit = *input.Unit
	}

	if input.City != nil {
		components.City = *input.City
	}

	if input.District != nil {
		components.District = *input.District
	}

	if input.State != nil {
		components.State = *input.State
	}

	if input.Country != nil {
		components.Country = *input.Country
	}

	if input.PostalCode != nil {
		components.PostalCode = *input.PostalCode
	}

	newComponents, err := e.Components.Update(components)
	if err != nil {
		return err
	}

	e.Components = newComponents
	return nil
}

// IsValid checks the validity of the address entity.
func (e *Entity) IsValid() error {
	err := validator.Struct(e)
	if err != nil {
		return err
	}

	err = e.Components.IsValid()
	if err != nil {
		return err
	}

	return err
}
