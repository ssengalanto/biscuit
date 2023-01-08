package gg_test

import (
	"testing"

	"github.com/ssengalanto/hex/pkg/gg"
	"github.com/stretchr/testify/require"
)

func TestPrepend(t *testing.T) {
	t.Parallel()

	n := []int{2, 3}
	n = gg.Prepend(n, 1)
	require.Equal(t, []int{1, 2, 3}, n)
}

func TestFindIndexOf(t *testing.T) {
	t.Parallel()

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
			idx := gg.FindIndexOf(tc.element, tc.predicate)
			require.Equal(t, tc.want, idx)
		})
	}
}

func TestItob(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input int
		want  bool
	}{
		{
			name:  "truthy",
			input: 23,
			want:  true,
		},
		{
			name:  "falsy",
			input: -1,
			want:  false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			b := gg.Itob(tc.input)
			require.Equal(t, tc.want, b)
		})
	}
}
