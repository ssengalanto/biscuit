//nolint:gochecknoglobals,gochecknoinits // unnecessary rules for this package
package validator

import (
	v "github.com/go-playground/validator/v10"
)

var validator *v.Validate

func init() {
	validator = v.New()
}

// Var validates a single variable using tag style validation.
func Var(field any, tag string) error {
	return validator.Var(field, tag)
}

// Struct validates a structs exposed fields, and automatically validates nested structs.
func Struct(s any) error {
	return validator.Struct(s)
}
