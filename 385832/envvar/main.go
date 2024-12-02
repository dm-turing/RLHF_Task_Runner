package envvar

import (
	"os"
)

// GetEnvWithDefault returns an environment variable value.
// If the variable is not set, it returns the provided default value.
func GetEnvWithDefault(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
