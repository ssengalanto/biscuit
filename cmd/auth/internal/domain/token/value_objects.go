package token

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ssengalanto/biscuit/pkg/validator"
)

const (
	GrantTypePassword = "grant-type-password"
)

var grantTypes = map[string]string{ //nolint:gochecknoglobals //intended
	GrantTypePassword: "password",
}

type JWT string

// NewJWT creates a new JWT token.
func NewJWT(p Payload, pk Base64RSAPrivateKey) (JWT, error) {
	key, err := pk.Parse()
	if err != nil {
		return "", err
	}

	claims := createClaims(p)

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrSigningTokenFailed, err)
	}

	return JWT(token), nil
}

// createClaims creates claims for JWT token.
func createClaims(p Payload) jwt.Claims {
	now := time.Now().UTC()

	c := make(jwt.MapClaims)

	c["sub"] = p.AccountID
	c["email"] = p.Email
	c["iss"] = p.Issuer
	c["aud"] = p.ClientID
	c["iat"] = now.Unix()
	c["nbf"] = now.Unix()
	c["exp"] = now.Add(p.Expiry).Unix()

	return c
}

// Validate checks the validity of the JWT token.
func (j JWT) Validate(pk Base64RSAPublicKey) (string, error) {
	token := j.String()

	key, err := pk.Parse()
	if err != nil {
		return "", err
	}

	pt, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("%w: %s", ErrInvalidSigningMethod, t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return "", ErrInvalidJWTToken
	}

	claims, ok := pt.Claims.(jwt.MapClaims)
	if !ok || !pt.Valid {
		return "", ErrInvalidJWTToken
	}

	sub := claims["sub"].(string) //nolint:errcheck //intentional panic
	return sub, nil
}

// String converts the JWTToken to string format.
func (j JWT) String() string {
	return string(j)
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
