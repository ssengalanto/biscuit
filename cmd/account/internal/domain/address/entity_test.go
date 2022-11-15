package address_test

import (
	"testing"

	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	entity := account.New()
	require.NotNilf(t, entity, "entity should not be nil")
}
