package mocks

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMockDB creates a new mock MySQL database for testing
func NewMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	// Create SQL mock
	sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create sql mock: %v", err)
	}

	// Create Gorm DB with MySQL dialect
	dialector := mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		sqlDB.Close()
		return nil, nil, fmt.Errorf("failed to open gorm db: %v", err)
	}

	return db, mock, nil
}

// BookColumns returns the standard columns for book queries
func BookColumns() []string {
	return []string{
		"id",
		"title",
		"author",
		"format",
		"file_path",
		"file_size",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}
