package json

import (
	"encoding/json"
	"net/http"

	"github.com/ssengalanto/biscuit/pkg/http/errors"
)

// EncodeError encodes HTTP error into JSON format.
func EncodeError(w http.ResponseWriter, err error) error {
	if err == nil {
		panic("EncodeError called for nil error")
	}

	httpError := errors.NewHTTPError(err)

	res, err := json.Marshal(httpError)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpError.Error.Code)
	w.Write(res) //nolint:errcheck // unnecessary

	return nil
}

// MustEncodeError encodes HTTP error into JSON format,
// if the encoding fails it will return an error.
func MustEncodeError(w http.ResponseWriter, err error) {
	if err = EncodeError(w, err); err != nil {
		panic(err)
	}
}
