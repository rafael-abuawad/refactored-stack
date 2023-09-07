package config

import (
	"fmt"
	"os"
)

var (
	// PORT returns the server listening port
	PORT = getEnv("PORT", "5000")
	// DB returns the name of the sqlite database
	DB = getEnv("DB", "refactored-stack.db")
	// TOKEN_PRIVATE_KEY returns the jwt token secret
	TOKEN_PRIVATE_KEY = getEnv("TOKEN_RPIVATE_KEY", "b63114f8-29fc-4c18-926d-2238e5de0d37")
	// TOKEN_EXPIRATION returns the jwt token expiration duration.
	TOKEN_EXPIRATION = getEnv("TOKEN_EXPIRATION", "10h")
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
