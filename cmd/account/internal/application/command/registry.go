package command

import (
	"github.com/mehdihadeli/go-mediatr"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
)

func RegisterHandlers(logger interfaces.Logger, repository account.Repository) {
	createAccountHandler := NewCreateAccountCommandHandler(logger, repository)
	err := mediatr.RegisterRequestHandler[*CreateAccountCommand, dto.CreateAccountResponseDto](createAccountHandler)
	if err != nil {
		logger.Fatal(err.Error(), nil)
	}
}
