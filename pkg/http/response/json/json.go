package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ssengalanto/potato-project/pkg/constants"
	"github.com/ssengalanto/potato-project/pkg/http/response"
)

// EncodeResponse encodes the payload to JSON format, if the encoding fails it will return an error.
func EncodeResponse(w http.ResponseWriter, statusCode int, payload any) error {
	res, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res) //nolint:errcheck // unnecessary

	response.Flush(w)

	return nil
}

// MustEncodeResponse encodes the payload to JSON format, if the encoding fails it will panic.
func MustEncodeResponse(w http.ResponseWriter, statusCode int, payload any) {
	if err := EncodeResponse(w, statusCode, payload); err != nil {
		panic(err)
	}
}

// DecodeRequest validates and decodes the JSON HTTP request body, if the validation fails it will return an error.
func DecodeRequest(w http.ResponseWriter, r *http.Request, dst any) error {
	r.Body = http.MaxBytesReader(w, r.Body, int64(constants.MaxHeaderBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("%w (at character %d)", ErrBadlyFormedJSON, syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return ErrBadlyFormedJSON

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("%w for field %q", ErrIncorrectJSONType, unmarshalTypeError.Field)
			}
			return fmt.Errorf("%w (at character %d)", ErrIncorrectJSONType, unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return ErrEmptyBody

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("%w %s", ErrUnknownField, fieldName)

		case err.Error() == "http: request body too large":
			return fmt.Errorf("%w, it must not be larger than %d bytes", ErrBodySizeExceeded, constants.MaxHeaderBytes)

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return ErrExtraneousBody
	}

	return nil
}
