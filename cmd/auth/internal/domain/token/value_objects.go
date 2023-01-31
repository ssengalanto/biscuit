package token

import (
	"github.com/ssengalanto/biscuit/pkg/validator"
)

const (
	GrantTypePassword = "grant-type-password"
)

var grantTypes = map[string]string{ //nolint:gochecknoglobals //intendedgit pu
	GrantTypePassword: "password",
}

type GrantType string

// IsValid checks the validity of the grant type.
func (g GrantType) IsValid() (bool, error) {
	err := validator.Var("GrantType", g, "required")
	if err != nil {
		return false, err
	}

	if _, ok := grantTypes[g.String()]; !ok {
		return false, ErrInvalidGrantType
	}

	return true, nil
}

// String converts GrantType to type string.
func (g GrantType) String() string {
	return string(g)
}
