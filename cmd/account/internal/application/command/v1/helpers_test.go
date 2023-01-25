package v1_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	acctmock "github.com/ssengalanto/biscuit/cmd/account/internal/mock"
	"github.com/ssengalanto/biscuit/pkg/mock"
)

func createDependencies(t *testing.T) (*mock.MockLogger, *acctmock.MockRepository, *acctmock.MockCache) {
	ctrl := gomock.NewController(t)
	logger := mock.NewMockLogger(ctrl)
	repository := acctmock.NewMockRepository(ctrl)
	cache := acctmock.NewMockCache(ctrl)
	return logger, repository, cache
}
