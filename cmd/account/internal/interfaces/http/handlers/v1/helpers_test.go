package v1_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ssengalanto/biscuit/pkg/mock"
)

func createDependencies(t *testing.T) (*mock.MockLogger, *mock.MockMediator) {
	ctrl := gomock.NewController(t)
	logger := mock.NewMockLogger(ctrl)
	mediator := mock.NewMockMediator(ctrl)
	return logger, mediator
}
