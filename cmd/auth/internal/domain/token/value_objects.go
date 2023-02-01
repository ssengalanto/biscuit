package token

import (
	"crypto/rsa"
	"encoding/base64"

	"github.com/golang-jwt/jwt"
	"github.com/ssengalanto/biscuit/pkg/validator"
)

const (
	GrantTypePassword = "grant-type-password"
)

var grantTypes = map[string]string{ //nolint:gochecknoglobals //intended
	GrantTypePassword: "password",
}

type Base64RSAPrivateKey string

// Parse decodes the base64 RSA private key and parse it to PKCS1 or PKCS8.
func (b Base64RSAPrivateKey) Parse() (*rsa.PrivateKey, error) {
	pk, err := base64.StdEncoding.DecodeString(b.String())
	if err != nil {
		return nil, ErrInvalidRSAPrivateKey
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(pk)
	if err != nil {
		return nil, ErrParseRSAPrivateKeyFailed
	}

	return key, nil
}

// String converts the Base64RSAPrivateKey to string format.
func (b Base64RSAPrivateKey) String() string {
	return string(b)
}

type Base64RSAPublicKey string

// Parse decodes the base64 RSA public key and parse it to PKCS1 or PKCS8.
func (b Base64RSAPublicKey) Parse() (*rsa.PublicKey, error) {
	pk, err := base64.StdEncoding.DecodeString(b.String())
	if err != nil {
		return nil, ErrInvalidRSAPublicKey
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(pk)
	if err != nil {
		return nil, ErrParseRSAPublicKeyFailed
	}

	return key, nil
}

// String converts the Base64RSAPublicKey to string format.
func (b Base64RSAPublicKey) String() string {
	return string(b)
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
