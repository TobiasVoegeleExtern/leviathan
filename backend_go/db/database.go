package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	var db *gorm.DB
	var err error
	maxRetries := 5
	for attempt := 1; attempt <= maxRetries; attempt++ {
		// Open the database connection using GORM's PostgreSQL driver
		db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
		if err == nil {
			// Successfully connected, break the loop
			break
		}

		// Log the error and retry after a delay
		log.Printf("Attempt %d to connect to the database failed: %v. Retrying...\n", attempt, err)
		time.Sleep(time.Duration(attempt) * time.Second) // Exponential backoff (retry after 1s, 2s, 3s, ...)
	}

	// If we still have an error after retries, log and return it
	if err != nil {
		log.Fatalf("Failed to connect to the database after %d attempts: %v", maxRetries, err)
	}

	// Optionally check if the database connection is successful
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Test the connection by pinging the database
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	// Return the GORM DB instance
	return db, nil
}
