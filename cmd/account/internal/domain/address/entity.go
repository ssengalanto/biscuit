package address

import (
	"github.com/google/uuid"
)

type Names struct {
	ShortName string `json:"shortName"`
	LongName  string `json:"longName"`
}

// Entity - Address Entity.
type Entity struct {
	ID         uuid.UUID  `json:"id"`
	PersonID   uuid.UUID  `json:"personId"`
	Components Components `json:"components"`
	Geometry   Geometry   `json:"geometry"`
}

type UpdateInput struct {
	Components Components
	Geometry   Geometry
}

// New creates a new address entity.
func New() Entity {
	return Entity{
		ID: uuid.New(),
	}
}

// Update checks the validity of the update address input
// and updates the address entity components and geometry fields.
func (e *Entity) Update(input UpdateInput) error {
	components, err := e.Components.Update(input.Components)
	if err != nil {
		return err
	}

	geometry, err := e.Geometry.Update(input.Geometry)
	if err != nil {
		return err
	}

	e.Components = components
	e.Geometry = geometry
	return nil
}
