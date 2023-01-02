package address

import (
	"github.com/google/uuid"
)

// Entity - address entity struct.
type Entity struct {
	ID         uuid.UUID  `json:"id"`
	PersonID   uuid.UUID  `json:"personId"`
	Components Components `json:"components"`
}

// New creates a new address entity.
func New() Entity {
	return Entity{
		ID: uuid.New(),
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
