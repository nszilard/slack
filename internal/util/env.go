package util

import "os"

// FallbackEnvString returns the environment value unless the variable is already set.
func FallbackEnvString(flag string, key string) string {
	if flag != "" {
		return flag
	}

	return os.Getenv(key)
}
