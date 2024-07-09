package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestNewJwtToken(t *testing.T) {
	assert := assert.New(t)

	_ = os.Setenv(jwtEnvKey, "other-secret-word")

	payload := jwt.MapClaims{
		"name": "John Doe",
	}

	token, err := NewJwtToken(payload, nil)

	assert.Nil(err, "err must be nil")
	assert.NotEmpty(token, "token must not empty")
}

func TestIsJwtTokenValid(t *testing.T) {
	assert := assert.New(t)

	payload := jwt.MapClaims{
		"name": "John Doe",
	}

	token, TokenErr := NewJwtToken(payload, nil)
	valid, err := IsJwtTokenValid(token)

	assert.Nil(TokenErr, "TokenErr must be nil")
	assert.True(valid, "valid must be true")
	assert.Nil(err, "err must be nil")
}

func TestIsJwtTokenValidFail(t *testing.T) {
	assert := assert.New(t)

	zeroDuration := -time.Second
	payload := jwt.MapClaims{
		"name": "John Doe",
	}

	token, TokenErr := NewJwtToken(payload, &zeroDuration)
	valid, err := IsJwtTokenValid(token)

	assert.Nil(TokenErr, "TokenErr must be nil")
	assert.False(valid, "valid must be false")
	assert.NotNil(err, "err must not be nil")
	assert.Equal("Token is expired", err.Error())
}

func TestGetJwtTokenMapClaims(t *testing.T) {
	assert := assert.New(t)

	payload := jwt.MapClaims{
		"name": "John Doe",
	}

	token, TokenErr := NewJwtToken(payload, nil)
	claims, err := GetJwtTokenMapClaims(token)
	delete(claims, "exp")
	delete(payload, "exp")

	assert.Nil(TokenErr, "TokenErr must be nil")
	assert.Equal(claims, payload)
	assert.Nil(err, "err must be nil")
}
