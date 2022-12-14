package json_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	httpjson "github.com/ssengalanto/potato-project/pkg/http/response/json"
	"github.com/stretchr/testify/require"
)

func TestEncodeResponse(t *testing.T) {
	user := newTestUser()

	w := httptest.NewRecorder()
	err := httpjson.EncodeResponse(w, http.StatusOK, user)
	require.NoError(t, err)

	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	require.Equal(t, res.Status, fmt.Sprintf("%d %s", http.StatusOK, http.StatusText(http.StatusOK)))
	require.Equal(
		t,
		string(data),
		fmt.Sprintf(`{"data":{"firstName":"%s","lastName":"%s"}}`, user.FirstName, user.LastName),
	)
}

func TestMustEncodeResponse(t *testing.T) {
	w := httptest.NewRecorder()
	require.NotPanics(t, func() {
		httpjson.EncodeResponse(w, http.StatusOK, nil) //nolint:errcheck //unnecessary
	})
}

func TestDecodeRequest(t *testing.T) {
	var match testUser
	user := newTestUser()

	r, err := json.Marshal(user)
	require.NoError(t, err)
	body := strings.NewReader(string(r))

	req := httptest.NewRequest(http.MethodGet, "/", body)

	w := httptest.NewRecorder()
	err = httpjson.DecodeRequest(w, req, &match)
	require.NoError(t, err)
	require.Equal(t, user, match)
}

type testUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func newTestUser() testUser {
	return testUser{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
	}
}
