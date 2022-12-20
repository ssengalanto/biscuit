package command

import (
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
	"github.com/ssengalanto/potato-project/pkg/mediatr"
)

func RegisterHandlers(logger interfaces.Logger, repository account.Repository, mediator *mediatr.Mediatr) {
	createAccountCommandHandler := NewCreateAccountCommandHandler(logger, repository)
	err := mediator.RegisterRequestHandler(createAccountCommandHandler)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}
}
