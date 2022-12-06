package pgsql_test

import (
	"fmt"
	"testing"

	"github.com/ssengalanto/potato-project/cmd/account/internal/infrastructure/persistence/pgsql"
	"github.com/stretchr/testify/require"
)

func TestMustBeValidAccountQuery(t *testing.T) {
	testCases := []struct {
		name   string
		assert func(t *testing.T)
	}{
		{
			name: "valid query",
			assert: func(t *testing.T) {
				require.NotPanics(t, func() {
					pgsql.MustBeValidAccountQuery(pgsql.QueryCreateAccount)
				})
			},
		},
		{
			name: "invalid query",
			assert: func(t *testing.T) {
				s := "invalid"
				errMsg := fmt.Sprintf("%s: `%s` doesn't exists in account queries", pgsql.ErrInvalidQuery.Error(), s)
				require.PanicsWithError(t, errMsg, func() {
					pgsql.MustBeValidAccountQuery(s)
				})
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.assert(t)
		})
	}
}
