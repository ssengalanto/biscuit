//nolint:gochecknoglobals,gochecknoinits,errorlint // unnecessary rules for this package
package validator

import (
	"bytes"
	"fmt"

	v "github.com/go-playground/validator/v10"
	"github.com/ssengalanto/biscuit/pkg/errors"
)

var validator *v.Validate

type fieldErr struct {
	key   string
	tag   string
	value any
}

func init() {
	validator = v.New()
}

// Var validates a single variable using tag style validation.
func Var(key string, field any, tag string) error {
	err := validator.Var(field, tag)
	if err != nil {
		return fmt.Errorf("%w: %s", errors.ErrInvalid, varErrMsg(err, key, field))
	}

	return nil
}

// Struct validates a structs exposed fields, and automatically validates nested structs.
func Struct(s any) error {
	err := validator.Struct(s)
	if err != nil {
		return fmt.Errorf("%w: %s", errors.ErrInvalid, structErrMsg(err))
	}

	return nil
}

// varErrMsg builds the error message for Var validator.
func varErrMsg(err error, key string, value any) string {
	var buf bytes.Buffer
	fieldErrs := fieldErrors(err)

	idx := 0
	for _, field := range fieldErrs {
		msg := fmt.Sprintf("`%s` field with value of `%v` failed on `%s` tag", key, value, field.tag)

		if idx >= 1 && idx != len(fieldErrs) {
			buf.WriteString(", " + msg)
			idx++
			continue
		}

		buf.WriteString(msg)
		idx++
	}

	return buf.String()
}

// structErrMsg builds the error message for Struct validator.
func structErrMsg(err error) string {
	var buf bytes.Buffer
	fieldErrs := fieldErrors(err)

	idx := 0
	for _, field := range fieldErrs {
		msg := fmt.Sprintf(
			"`%s` field with value of `%v` failed on `%s` tag validation",
			field.key,
			field.value,
			field.tag,
		)

		if idx >= 1 && idx != len(fieldErrs) {
			buf.WriteString(", " + msg)
			idx++
			continue
		}

		buf.WriteString(msg)
		idx++
	}

	return buf.String()
}

// fieldErrors iterates through validation errors and stores key, tag and value inside a map.
func fieldErrors(err error) map[string]fieldErr {
	fieldErrs := make(map[string]fieldErr)

	for _, err := range err.(v.ValidationErrors) {
		if _, ok := fieldErrs[err.StructNamespace()]; !ok {
			fieldErrs[err.StructNamespace()] = fieldErr{
				key:   err.StructNamespace(),
				tag:   err.Tag(),
				value: err.Value(),
			}
		}
	}

	return fieldErrs
}
