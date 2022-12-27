package v1

import (
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
)

// GetAccountQuery contains required fields for account retrieval.
type GetAccountQuery struct {
	ID string `json:"id"`
}

// NewGetAccountQuery creates a new query for account retrieval.
func NewGetAccountQuery(input dto.GetAccountRequestDto) *GetAccountQuery {
	return &GetAccountQuery{
		ID: input.ID,
	}
}
