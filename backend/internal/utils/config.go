package utils

import (
	"os"
)

// Config  ...
type Config struct {
	ServerPort string
	LogLevel   string
}

// LoadConfig ....
func LoadConfig() *Config {
	config := &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		LogLevel:   getEnv("LOG_LEVEL", "info"),
	}
	return config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
