package utils

import (
	"strings"

	"go.uber.org/zap"
)

// NewLogger ...
func NewLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction() // Use zap.NewDevelopment() for debugging
	if err != nil {
		return nil, err
	}
	defer logger.Sync()
	return logger, nil
}

// NewCustomLogger ...
func NewCustomLogger() (*zap.Logger, error) {
	logLevel := strings.ToLower(getEnv("LOG_LEVEL", "info"))
	var level zap.AtomicLevel

	switch strings.ToLower(logLevel) {
	case "debug":
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "error":
		level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		level = zap.NewAtomicLevelAt(zap.InfoLevel) // Default INFO
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = level

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
