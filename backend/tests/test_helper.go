package tests

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/zven/bookpavilion/config"
	"github.com/zven/bookpavilion/mocks"
	"gorm.io/gorm"
)

var (
	testDB *gorm.DB
	mock   sqlmock.Sqlmock
)

// SetupTestDB initializes a mock MySQL database for testing
func SetupTestDB(t *testing.T) sqlmock.Sqlmock {
	var err error
	testDB, mock, err = mocks.NewMockDB()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}

	// Set the mock database in config
	config.SetDB(testDB)

	return mock
}

// SetupTestEnvironment prepares the test environment
func SetupTestEnvironment(t *testing.T) sqlmock.Sqlmock {
	// Reset config
	config.Reset()

	// Create temporary test uploads directory
	testUploadDir := GetUploadDir()
	err := os.MkdirAll(testUploadDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test upload directory: %v", err)
	}

	// Set the test upload directory in config
	config.SetUploadDir(testUploadDir)

	// Setup mock database
	return SetupTestDB(t)
}

// CleanupTestEnvironment cleans up the test environment
func CleanupTestEnvironment(t *testing.T) {
	// Clean up test uploads directory
	testUploadDir := config.GetUploadDir()
	err := os.RemoveAll(testUploadDir)
	if err != nil {
		t.Errorf("Failed to cleanup test upload directory: %v", err)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %v", err)
	}

	// Reset config
	config.Reset()
}
