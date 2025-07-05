package jwtutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")

	token, err := GenerateToken(123)
	assert.NoError(t, err)

	userID, err := ParseToken(token)
	assert.NoError(t, err)
	assert.Equal(t, 123, userID)
}
