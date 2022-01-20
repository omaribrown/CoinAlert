package envVariables

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv_Get(t *testing.T) {
	t.Run("it should get environment variables from .env when APP_ENV is set to isLocal", func(t *testing.T) {
		// BEFORE
		os.Unsetenv("APP_ENV")


		// ARRANGE
		os.Setenv("APP_ENV", "local")
		os.Setenv("TEST_VAR", "some-var")
		env, err := New(Props{
			DotEnvPath: ".env",
		})

		assert.NoError(t, err)

		// ACT
		result := env.Get("TEST_VAR")

		// ASSERT
		assert.Equal(t, "this-var", result)
	})

	t.Run("it should get environment variables from .env when APP_ENV is set to production", func(t *testing.T) {
		// BEFORE
		os.Unsetenv("APP_ENV")


		// ARRANGE
		os.Setenv("APP_ENV", "production")
		os.Setenv("TEST_VAR", "some-var")
		env, err := New(Props{
			DotEnvPath: ".env",
		})

		assert.NoError(t, err)

		// ACT
		result := env.Get("TEST_VAR")

		// ASSERT
		assert.Equal(t, "some-var", result)
	})
}