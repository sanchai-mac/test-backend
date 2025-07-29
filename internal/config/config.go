package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	_ "time/tzdata"
)

// Configuration ...
type Configuration struct {
	CustomerDB     string `env:"CUSTORMER_DB" envDefault:"postgres://postgres:password123@localhost:11432/test_customer_db?sslmode=disable"`
	Port           string
	FilesDirectory string
}

// NewConfiguration ...
func NewConfiguration() *Configuration {
	// Try to load .env file if it exists
	if err := loadEnvFile(); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	cfg := &Configuration{
		//		CustomerDB: viper.GetString("CUSTORMER_DB"),
	}
	cfg.Port = os.Getenv("PORT")
	cfg.CustomerDB = os.Getenv("CUSTORMER_DB")

	return cfg
}

// loadEnvFile tries to find and load .env file
func loadEnvFile() error {
	// Get current working directory
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Search for .env file starting from current directory up to root
	dir := wd
	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			fmt.Printf("Loading .env file from: %s\n", envPath)
			return godotenv.Load(envPath)
		}

		// Go up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root directory
			break
		}
		dir = parent
	}

	return fmt.Errorf(".env file not found")
}
