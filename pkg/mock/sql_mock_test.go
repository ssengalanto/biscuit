package mock_test

import (
	"testing"

	"github.com/ssengalanto/biscuit/pkg/mock"
	"github.com/stretchr/testify/require"
)

func TestNewSqlDb(t *testing.T) {
	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	require.NotNil(t, db)
	require.NotNil(t, dbmock)
}
