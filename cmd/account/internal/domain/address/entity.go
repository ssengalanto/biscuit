package address

import (
	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/pkg/validator"
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

// Update checks the validity of the update address input
// and updates the address entity components and geometry fields.
func (e *Entity) Update(input Components) error {
	components, err := e.Components.Update(input)
	if err != nil {
		return err
	}

	e.Components = components
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
