package utils

import "os"

// ga usah dipake gpp si cuma enak aja klo ada
func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}