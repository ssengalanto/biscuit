package token_test

import (
	"crypto/rsa"
	"reflect"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ssengalanto/biscuit/cmd/auth/internal/domain/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewJWT(t *testing.T) {
	p := newTokenPayload()
	pk := token.Base64RSAPrivateKey(newBase64RSAPrivateKey())

	t.Run("it should create a new JWT token", func(t *testing.T) {
		tk, err := token.NewJWT(p, pk)
		assert.NotEmpty(t, tk)
		require.NoError(t, err)
	})

	t.Run("it should fail to create a new JWT Token", func(t *testing.T) {
		ipk := token.Base64RSAPrivateKey(newBase64RSAPublicKey())
		tk, err := token.NewJWT(p, ipk)
		assert.Empty(t, tk)
		require.Error(t, err)
	})
}

func TestJWT_Validate(t *testing.T) {
	p := newTokenPayload()
	pvk := token.Base64RSAPrivateKey(newBase64RSAPrivateKey())
	pbk := token.Base64RSAPublicKey(newBase64RSAPublicKey())

	t.Run("it should validate the JWT token successfully", func(t *testing.T) {
		tk, err := token.NewJWT(p, pvk)
		require.NoError(t, err)

		res, err := tk.Validate(pbk)
		assert.Equal(t, p.AccountID, res)
		require.NoError(t, err)
	})

	t.Run("it should fail to validate the JWT token", func(t *testing.T) {
		tk, err := token.NewJWT(p, pvk)
		require.NoError(t, err)

		res, err := tk.Validate("")
		assert.Empty(t, res)
		require.Error(t, err)
	})
}

func TestJWT_String(t *testing.T) {
	t.Run("it should convert JWT to string", func(t *testing.T) {
		jwt := token.JWT("token").String()
		kind := reflect.TypeOf(jwt).String()
		require.Equal(t, "string", kind)
	})
}

func TestGrantType_IsValid(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, result bool, err error)
	}{
		{
			name:    "it should be a valid grant type",
			payload: token.GrantTypePassword,
			assert: func(t *testing.T, result bool, err error) {
				assert.True(t, result)
				require.NoError(t, err)
			},
		},
		{
			name:    "it should fail the validation due to empty value",
			payload: "",
			assert: func(t *testing.T, result bool, err error) {
				assert.False(t, result)
				require.Error(t, err)
			},
		},
		{
			name:    "it should be an invalid grant type",
			payload: "invalid-grant-type",
			assert: func(t *testing.T, result bool, err error) {
				assert.False(t, result)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gt := token.GrantType(tc.payload)
			ok, err := gt.IsValid()
			tc.assert(t, ok, err)
		})
	}
}

func TestGrantType_String(t *testing.T) {
	t.Run("it should convert email to string", func(t *testing.T) {
		gt := token.GrantType("password").String()
		kind := reflect.TypeOf(gt).String()
		require.Equal(t, "string", kind)
	})
}

func TestBase64RSAPrivateKey_Parse(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, key *rsa.PrivateKey, err error)
	}{
		{
			name:    "it should parse RSA private key successfully",
			payload: newBase64RSAPrivateKey(),
			assert: func(t *testing.T, key *rsa.PrivateKey, err error) {
				assert.NotNil(t, key)
				require.NoError(t, err)
			},
		},
		{
			name:    "it should fail to decode base64 RSA private key",
			payload: "invalid-base64-rsa-private-key",
			assert: func(t *testing.T, key *rsa.PrivateKey, err error) {
				assert.Nil(t, key)
				require.Error(t, err)
			},
		},
		{
			name:    "it should fail to parse RSA private key from PEM",
			payload: newBase64RSAPublicKey(),
			assert: func(t *testing.T, key *rsa.PrivateKey, err error) {
				assert.Nil(t, key)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pk := token.Base64RSAPrivateKey(tc.payload)
			key, err := pk.Parse()
			tc.assert(t, key, err)
		})
	}
}

func TestBase64RSAPrivateKey_String(t *testing.T) {
	t.Run("it should convert base64 RSA public key to string", func(t *testing.T) {
		pk := newBase64RSAPrivateKey()
		kind := reflect.TypeOf(pk).String()
		require.Equal(t, "string", kind)
	})
}

func TestBase64RSAPublicKey_Parse(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		assert  func(t *testing.T, key *rsa.PublicKey, err error)
	}{
		{
			name:    "it should parse RSA public key successfully",
			payload: newBase64RSAPublicKey(),
			assert: func(t *testing.T, key *rsa.PublicKey, err error) {
				assert.NotNil(t, key)
				require.NoError(t, err)
			},
		},
		{
			name:    "it should fail to decode base64 RSA public key",
			payload: "invalid-base64-rsa-private-key",
			assert: func(t *testing.T, key *rsa.PublicKey, err error) {
				assert.Nil(t, key)
				require.Error(t, err)
			},
		},
		{
			name:    "it should fail to parse RSA public key from PEM",
			payload: newBase64RSAPrivateKey(),
			assert: func(t *testing.T, key *rsa.PublicKey, err error) {
				assert.Nil(t, key)
				require.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pk := token.Base64RSAPublicKey(tc.payload)
			key, err := pk.Parse()
			tc.assert(t, key, err)
		})
	}
}

func TestBase64RSAPublicKey_String(t *testing.T) {
	t.Run("it should convert base64 RSA private key to string", func(t *testing.T) {
		pk := newBase64RSAPublicKey()
		kind := reflect.TypeOf(pk).String()
		require.Equal(t, "string", kind)
	})
}

func newTokenPayload() token.Payload {
	return token.Payload{
		AccountID: gofakeit.UUID(),
		ClientID:  gofakeit.UUID(),
		Issuer:    gofakeit.Word(),
		ExpiresIn: 15 * time.Minute,
	}
}

func newBase64RSAPrivateKey() string {
	return "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDV3dJQkFBS0JnUUNBeC9mVmgyUmtkSDRiOVZscHpQWFpreGcxYnptQ0pQeVpNN0pTQldNRXAwbU1SZE1ICkJRZDdvUzZCSWlqR3htTWZzT3JBajhFWjg4d0MxbDVSdEg5VFpwcE9sYzZ5dHRYRHJSZjhVQlRSVG55dUNoalkKdHhUVTQ5cE1HTVhNVkw2OVVZR3BGeCtENWRGZDlmVWZVODY2TmNUWE8rT1hyTlhtWGxxL2pTTWxyd0lEQVFBQgpBb0dBZWdhQkZLaUUvUmJSQkFiNFlXTWZ0YmxHb0NNekY5bWFMRVNxL0VNMGJ3MWdpSFVGSDhxcEs0RXdBcFp1Cmt1TWF1OFcwdXgrNzlxNW5LbTBiMUVtMnR0dTlyazVSdWhoZ3ZDZHZsalJUTlZDcDV2TEVWN0psOVl2Z2d1aEUKQzFLMUF5cWtQREE0eXY5WFdlajY2M3JhSWZ0bHpLRmloaXg1a1JwWkxQQUdQaEVDUVFEbEQ5S2U4Q3hFQVpWTAp1UkhVWXhwcTZSZ0o0eEphL2FoMHo1c0xVNTE1TU1peWpraGpJY3lhU082MDV4ZEhXd2VvZ0ZIdGpiUkZGTmpzCmV1Y2ZveXYzQWtFQWorMFRtS1Y2SGl6T3VmaTFwRmtjc1lJcWczQlZFMnp1NE1XTDhYSlJFQWdya3hmR2lzM0QKUkI2TmRDMlNCaHJmS2F5ZTVzcVNhbWZxSHJIUUhzTzJDUUpBSUQ4TC9ZZitFMHpOd2EwNkQxWXNQK1MwbDUrNQowOGxsejV2eVRiUGx0VXZpMVJBbXJKM3plYnpPcmZUaVdBOCtrc0FOeUkxc1ZWVkwvRzZJM3ZGUG5RSkFWa3pnCjREbnhKS0RYZ0huYWFPYXFKdUlYSGVOQWtEcFVibURsemV3dklUN1U2Z2xxbXBaUXpNckpKTzJpVHBqVVVZZloKYkNmeGJXNUwyd1hoOW1DQ0NRSkFKRzNUemJKdkNRajBmczB5S1dlOEphblhjTUlMQ21nTWhpQU9RK3JSNGtnTgpVMk1EaUoydmxhMmV6dDBJK1l5ZGRoNWZCbmE1d1d2WGpDMVRPWnE2Ymc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQ==" //nolint:lll //unnecessary
}

func newBase64RSAPublicKey() string {
	return "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FDQXgvZlZoMlJrZEg0YjlWbHB6UFhaa3hnMQpiem1DSlB5Wk03SlNCV01FcDBtTVJkTUhCUWQ3b1M2Qklpakd4bU1mc09yQWo4RVo4OHdDMWw1UnRIOVRacHBPCmxjNnl0dFhEclJmOFVCVFJUbnl1Q2hqWXR4VFU0OXBNR01YTVZMNjlVWUdwRngrRDVkRmQ5ZlVmVTg2Nk5jVFgKTytPWHJOWG1YbHEvalNNbHJ3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQ==" //nolint:lll //unnecessary
}
