package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tedysaputro/book-catalog-with-go/src/author"
	"github.com/tedysaputro/book-catalog-with-go/src/book"
	"github.com/tedysaputro/book-catalog-with-go/src/category"
	"github.com/tedysaputro/book-catalog-with-go/src/publisher"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnvOrDefault("DB_HOST", "localhost"),
		getEnvOrDefault("DB_USER", "postgres"),
		getEnvOrDefault("DB_PASSWORD", "postgres"),
		getEnvOrDefault("DB_NAME", "book_catalog"),
		getEnvOrDefault("DB_PORT", "5432"),
	)

	// Configure GORM logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Info,   // Log level (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: false,        // Include not found error
			Colorful:                  true,         // Enable color
		},
	)

	// Open database connection with logger
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Set the database instance for the models
	author.SetDB(db)
	publisher.SetDB(db)
	category.SetDB(db)
	book.SetDB(db)
	DB = db

	// Auto migrate the database
	err = DB.AutoMigrate(&author.Author{}, &publisher.Publisher{}, &category.Category{}, &book.Book{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected and migrated successfully")
}

// getEnvOrDefault returns environment variable value or default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
