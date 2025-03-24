package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ShivaayAPIKey string
	ServerPort    string
	Temperature   float64
	TopP          float64
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	temperature, err := strconv.ParseFloat(getEnvOrDefault("TEMPERATURE", "0.7"), 64)
	if err != nil {
		return nil, fmt.Errorf("invalid temperature value: %w", err)
	}

	topP, err := strconv.ParseFloat(getEnvOrDefault("TOP_P", "1.0"), 64)
	if err != nil {
		return nil, fmt.Errorf("invalid top_p value: %w", err)
	}

	config := &Config{
		ShivaayAPIKey: os.Getenv("SHIVAAY_API_KEY"),
		ServerPort:    getEnvOrDefault("SERVER_PORT", "8080"),
		Temperature:   temperature,
		TopP:          topP,
	}

	if config.ShivaayAPIKey == "" {
		return nil, fmt.Errorf("SHIVAAY_API_KEY is required")
	}

	return config, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
