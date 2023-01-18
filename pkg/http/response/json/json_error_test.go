package json_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ssengalanto/biscuit/pkg/errors"
	httpjson "github.com/ssengalanto/biscuit/pkg/http/response/json"
	"github.com/stretchr/testify/require"
)

func TestEncodeError(t *testing.T) {
	testCases := []struct {
		name      string
		apperror  error
		httperror int
	}{
		{
			name:      "bad request",
			apperror:  errors.ErrInvalid,
			httperror: http.StatusBadRequest,
		},
		{
			name:      "unauthorized",
			apperror:  errors.ErrUnauthorized,
			httperror: http.StatusUnauthorized,
		},
		{
			name:      "forbidden",
			apperror:  errors.ErrForbidden,
			httperror: http.StatusForbidden,
		},
		{
			name:      "not found",
			apperror:  errors.ErrNotFound,
			httperror: http.StatusNotFound,
		},
		{
			name:      "request timeout",
			apperror:  errors.ErrTimeout,
			httperror: http.StatusRequestTimeout,
		},
		{
			name:      "service unavailable",
			apperror:  errors.ErrTemporaryDisabled,
			httperror: http.StatusServiceUnavailable,
		},
		{
			name:      "internal server error",
			apperror:  errors.ErrInternal,
			httperror: http.StatusInternalServerError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := httpjson.EncodeError(w, tc.apperror)
			require.NoError(t, err)

			res := w.Result()
			defer res.Body.Close()

			data, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			require.Equal(t, res.Status, fmt.Sprintf("%d %s", tc.httperror, http.StatusText(tc.httperror)))
			require.Equal(
				t,
				string(data),
				fmt.Sprintf(
					`{"error":{"code":%d,"message":"%s","reason":"%s"}}`,
					tc.httperror,
					http.StatusText(tc.httperror),
					tc.apperror.Error(),
				),
			)
		})
	}
}

func TestMustEncodeError(t *testing.T) {
	apperror := errors.ErrInvalid
	w := httptest.NewRecorder()
	require.NotPanics(t, func() {
		httpjson.MustEncodeError(w, apperror)
	})
}
