package envvar

import (
	"fmt"
	"os"
	"strings"
)

// GetEnv retrieves an environment variable using os.Getenv.
//
// If the environment variable is not set, GetEnv returns the defaultValue.
// If the key is empty, it returns an error.
func GetEnv(key string, defaultValue string) (string, error) {
	if strings.TrimSpace(key) == "" {
		return "", ErrEmptyKey
	}

	if value := os.Getenv(key); value != "" {
		return value, nil
	}
	return defaultValue, nil
}

// ErrEmptyKey is returned when the key is empty.
var ErrEmptyKey = fmt.Errorf("empty key")
