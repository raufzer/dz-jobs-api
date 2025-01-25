package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	t.Run("Valid environment variable", func(t *testing.T) {
		os.Setenv("TEST_KEY", "test_value")
		defer os.Unsetenv("TEST_KEY")

		val, err := GetEnv("TEST_KEY", "string")
		assert.NoError(t, err)
		assert.Equal(t, "test_value", val.(string))
	})

	t.Run("Missing environment variable", func(t *testing.T) {
		// Test for a missing environment variable
		val, err := GetEnv("NON_EXISTENT_KEY", "string")
		assert.Nil(t, val, "Value should be nil for a missing environment variable")
		assert.Error(t, err, "An error is expected for a missing environment variable")
		assert.Contains(t, err.Error(), "environment variable NON_EXISTENT_KEY not set")
	})

	t.Run("Invalid duration format", func(t *testing.T) {
		os.Setenv("TEST_DURATION", "invalid")
		defer os.Unsetenv("TEST_DURATION")

		val, err := GetEnv("TEST_DURATION", "duration")
		assert.Nil(t, val, "Value should be nil for an invalid duration format")
		assert.Error(t, err, "An error is expected for an invalid duration format")
		assert.Contains(t, err.Error(), "failed to parse duration")
	})

	t.Run("Unsupported value type", func(t *testing.T) {
		os.Setenv("TEST_KEY", "test_value")
		defer os.Unsetenv("TEST_KEY")

		val, err := GetEnv("TEST_KEY", "unsupported")
		assert.Nil(t, val, "Value should be nil for an unsupported value type")
		assert.Error(t, err, "An error is expected for an unsupported value type")
		assert.Contains(t, err.Error(), "unsupported value type")
	})
}


