package token

import "fmt"

// Errors used by the token entity.
var (
	// ErrInvalidGrantType is returned when the grant type is invalid.
	ErrInvalidGrantType = fmt.Errorf("invalid grant type")
	// ErrInvalidRSAPrivateKey is returned when the rsa private key is invalid.
	ErrInvalidRSAPrivateKey = fmt.Errorf("invalid rsa private key")
	// ErrParseRSAPrivateKeyFailed is returned when parsing rsa private key failed.
	ErrParseRSAPrivateKeyFailed = fmt.Errorf("parsing rsa private key failed")
	// ErrInvalidRSAPublicKey is returned when the rsa public key is invalid.
	ErrInvalidRSAPublicKey = fmt.Errorf("invalid rsa public key")
	// ErrParseRSAPublicKeyFailed is returned when parsing rsa public key failed.
	ErrParseRSAPublicKeyFailed = fmt.Errorf("parsing rsa public failed")
	// ErrInvalidJWTToken is returned when the JWT token is invalid.
	ErrInvalidJWTToken = fmt.Errorf("invalid jwt token")
	// ErrInvalidSigningMethod is returned when unexpected method is used for signing JWT token.
	ErrInvalidSigningMethod = fmt.Errorf("unexpected method")
	// ErrSigningTokenFailed is returned when signing token failed.
	ErrSigningTokenFailed = fmt.Errorf("signing token failed")
)
