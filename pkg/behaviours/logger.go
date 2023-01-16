package behaviours

import (
	"context"
	"fmt"

	"github.com/ssengalanto/biscuit/pkg/interfaces"
	"github.com/ssengalanto/midt"
)

type LoggerBehaviour struct {
	log interfaces.Logger
}

func NewLoggerBehaviour(logger interfaces.Logger) *LoggerBehaviour {
	return &LoggerBehaviour{
		log: logger,
	}
}

func (l *LoggerBehaviour) Handle(
	ctx context.Context,
	request any,
	next midt.RequestHandlerFunc,
) (any, error) {
	l.log.Info(fmt.Sprintf("executing %T", request), map[string]any{"request": request})

	res, err := next()
	if err != nil {
		return nil, err
	}

	return res, nil
}
