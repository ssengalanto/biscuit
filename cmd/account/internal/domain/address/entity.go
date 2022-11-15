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

// New creates a new address entity.
func New() Entity {
	return Entity{
		ID: uuid.New(),
	}
}
