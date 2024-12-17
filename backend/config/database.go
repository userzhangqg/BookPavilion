package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/zven/bookpavilion/models"
)

// InitDB initializes the database connection
func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "root"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "bookpavilion"),
		getEnv("DB_CHARSET", "utf8mb4"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Auto Migrate the schema
	if err := db.AutoMigrate(&models.Book{}); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	SetDB(db)
	log.Println("Database connection established")
	return nil
}
