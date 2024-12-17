package config

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/gorm"
)

// Config holds all configuration for the application
type Config struct {
	DB        *gorm.DB
	UploadDir string
}

var (
	appConfig Config
	once      sync.Once
)

// LoadConfig initializes the application configuration
func LoadConfig() error {
	var err error
	once.Do(func() {
		// Set up upload directory
		uploadDir := getEnv("UPLOAD_DIR", "./uploads")
		if err = os.MkdirAll(uploadDir, 0755); err != nil {
			err = fmt.Errorf("failed to create upload directory: %v", err)
			return
		}
		appConfig.UploadDir = uploadDir
	})
	return err
}

// SetDB sets the database instance
func SetDB(db *gorm.DB) {
	appConfig.DB = db
}

// SetUploadDir sets the upload directory
func SetUploadDir(dir string) {
	appConfig.UploadDir = dir
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return appConfig.DB
}

// GetUploadDir returns the upload directory path
func GetUploadDir() string {
	return appConfig.UploadDir
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Reset resets the configuration (useful for testing)
func Reset() {
	appConfig = Config{}
	once = sync.Once{}
}
