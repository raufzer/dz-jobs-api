package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	secret := "supersecretkey"

	t.Run("Valid Token Generation", func(t *testing.T) {
		token, err := GenerateToken("user123", time.Hour, "access", "admin", secret)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		// Try to parse and validate the generated token
		claims, err := ValidateToken(token, secret, "access")
		assert.NoError(t, err)
		assert.Equal(t, "user123", claims.ID)
		assert.Equal(t, "admin", claims.Role)
		assert.Equal(t, "access", claims.Purpose)
	})

	t.Run("Missing Role for Access Token", func(t *testing.T) {
		token, err := GenerateToken("user123", time.Hour, "access", "", secret)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		claims, err := ValidateToken(token, secret, "access")
		assert.NoError(t, err)
		assert.Equal(t, "user123", claims.ID)
		assert.Empty(t, claims.Role)
	})

	t.Run("Invalid Purpose", func(t *testing.T) {
		token, err := GenerateToken("user123", time.Hour, "access", "admin", secret)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		_, err = ValidateToken(token, secret, "reset")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "token purpose mismatch")
	})

	t.Run("Invalid Token Signature", func(t *testing.T) {
		token, err := GenerateToken("user123", time.Hour, "access", "admin", secret)
		assert.NoError(t, err)

		_, err = ValidateToken(token, "wrongsecretkey", "access")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid token")
	})
}
