package v1_test

import (
	"github.com/golang/mock/gomock"
	acctmock "github.com/ssengalanto/biscuit/cmd/account/internal/mock"
	"github.com/ssengalanto/biscuit/pkg/mock"
)

func createDepedencies(ctrl *gomock.Controller) (*mock.MockLogger, *acctmock.MockRepository, *acctmock.MockCache) {
	logger := mock.NewMockLogger(ctrl)
	repository := acctmock.NewMockRepository(ctrl)
	cache := acctmock.NewMockCache(ctrl)
	return logger, repository, cache
}
