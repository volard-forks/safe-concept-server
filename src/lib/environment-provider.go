package lib

import (
	"fmt"
	"os"
)

// noEnvPanic is an internal helper function that generates a panic when an environment
// variable is missing.
func noEnvPanic(env string) {
	panic(fmt.Sprintf("environment_provider: %v env reading", env))
}

// GetEnv retrieves an environment variable by name.
// Enforces that all required environment variables must be set.
//
// Parameters:
//   - name: The name of the environment variable to retrieve
//
// Returns:
//   - The value of the environment variable if it exists
func GetEnv(name string) string {
	env := os.Getenv(name)
	if env == "" {
		noEnvPanic(name)
	}

	return env
}
