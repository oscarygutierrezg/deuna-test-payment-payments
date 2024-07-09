package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	assert := assert.New(t)
	const password = "pass"

	saltSrt, hashStr := EncryptPassword(password)

	assert.NotEmpty(saltSrt)
	assert.NotEmpty(hashStr)
	assert.NotEqual(password, hashStr)
}

func TestVerifyPassword(t *testing.T) {
	assert := assert.New(t)
	const password = "pass"

	saltSrt, hashStr := EncryptPassword(password)
	ok := VerifyPassword(password, saltSrt, hashStr)

	assert.NotEmpty(saltSrt)
	assert.NotEmpty(hashStr)
	assert.NotEqual(password, hashStr)
	assert.True(ok)
}

func TestVerifyPasswordFail(t *testing.T) {
	assert := assert.New(t)
	const password = "pass"

	saltSrt, hashStr := EncryptPassword(password)
	ok := VerifyPassword(password, saltSrt, hashStr+"fail")

	assert.NotEmpty(saltSrt)
	assert.NotEmpty(hashStr)
	assert.NotEqual(password, hashStr)
	assert.False(ok)
}
