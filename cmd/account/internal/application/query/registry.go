package query

import (
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
	"github.com/ssengalanto/potato-project/pkg/mediatr"
)

// RegisterHandlers registers all the query handlers in the query registry.
func RegisterHandlers(logger interfaces.Logger, repository account.Repository, mediator *mediatr.Mediatr) {
	getAccountQueryHandler := NewGetAccountQueryHandler(logger, repository)
	err := mediator.RegisterRequestHandler(getAccountQueryHandler)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}
}
