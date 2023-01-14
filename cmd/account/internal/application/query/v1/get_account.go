package v1

import (
	"github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto"
)

// GetAccountQuery contains required fields for account retrieval.
type GetAccountQuery struct {
	ID string `json:"id"`
}

// NewGetAccountQuery creates a new query for account retrieval.
func NewGetAccountQuery(input dto.GetAccountRequest) *GetAccountQuery {
	return &GetAccountQuery{
		ID: input.ID,
	}
}
