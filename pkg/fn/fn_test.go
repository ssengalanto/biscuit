package fn_test

import (
	"github.com/ssengalanto/potato-project/pkg/fn"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrepend(t *testing.T) {
	t.Parallel()

	n := []int{2, 3}
	n = fn.Prepend(n, 1)
	require.Equal(t, []int{1, 2, 3}, n)
}

func TestFindIndexOf(t *testing.T) {
	testCases := []struct {
		name      string
		element   []int
		predicate func(n int) bool
		want      int
	}{
		{
			name:    "element found",
			element: []int{1, 2, 3},
			predicate: func(n int) bool {
				return n == 2
			},
			want: 1,
		},
		{
			name:    "element not found",
			element: []int{1, 2, 3},
			predicate: func(n int) bool {
				return n == 0
			},
			want: -1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			idx := fn.FindIndexOf(tc.element, tc.predicate)
			require.Equal(t, tc.want, idx)
		})
	}
}
