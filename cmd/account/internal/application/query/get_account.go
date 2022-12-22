package query

import (
	"fmt"

	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
)

// GetAccountQuery contains required fields for account retrieval, satisfies mediatr.Request.
type GetAccountQuery struct {
	ID string `json:"id"`
}

func (c *GetAccountQuery) Name() string {
	return fmt.Sprintf("%T", &GetAccountQuery{})
}

// NewGetAccountQuery creates a new query for account retrieval.
func NewGetAccountQuery(input dto.GetAccountRequestDto) *GetAccountQuery {
	return &GetAccountQuery{
		ID: input.ID,
	}
}
