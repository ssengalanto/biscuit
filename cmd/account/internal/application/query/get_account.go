package query

import (
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
)

type GetAccountQuery struct {
	ID string `json:"id"`
}

func (c *GetAccountQuery) Topic() string {
	return QueryGetAccount
}

func NewGetAccountQuery(input dto.GetAccountRequestDto) *GetAccountQuery {
	return &GetAccountQuery{
		ID: input.ID,
	}
}
