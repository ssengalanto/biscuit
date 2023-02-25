package v1

import (
	dtov1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto/v1"
)

// GetAccountQuery contains required fields for account retrieval.
type GetAccountQuery struct {
	ID string `json:"id"`
}

// NewGetAccountQuery creates a new query for account retrieval.
func NewGetAccountQuery(input dtov1.GetAccountRequest) *GetAccountQuery {
	return &GetAccountQuery{
		ID: input.ID,
	}
}
