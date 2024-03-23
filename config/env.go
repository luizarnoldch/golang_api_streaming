package config

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func sanityCheck() {
	requiredEnvVars := []string{
		"API_HOST",
		"API_PORT",
		"ENV",
		"STREAM_TABLE",
	}

	for _, envVar := range requiredEnvVars {
		if value := os.Getenv(envVar); value == "" {
			log.Fatalf("Environment variable %s not defined. Terminating application...", envVar)
		}
	}
}

func LoadConfig(filePath string) (*CONFIG, error) {
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatalf("Error loading .en %v", err)
		return nil, fmt.Errorf("error loading .env %v", err)
	}

	sanityCheck()

	return &CONFIG{
		MICRO: MICRO{
			API: API{
				HOST: os.Getenv("API_HOST"),
				PORT: os.Getenv("API_PORT"),
			},
			DB: DB{
				STREAM_DYNAMODB: STREAM_DYNAMODB{
					TABLE_NAME: os.Getenv("STREAM_TABLE"),
				},
			},
		},
		ENV: os.Getenv("ENV"),
	}, nil
}
