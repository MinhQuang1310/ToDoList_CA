package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the database connection.
// It loads the environment variables from the .env file,
// constructs the database connection string, and establishes
// the database connection.
// It returns the initialized database connection.
func InitDB() *gorm.DB {
	// Load environment variables from .env file.
	err := godotenv.Load()

	if err != nil {
		// If the .env file is not found, panic with an error message.
		panic("Error loading .env file")
	}

	// Get the database host, user, name, SSL mode, and password from environment variables.
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	dbPassword := os.Getenv("DB_PASSWORD")

	// Construct the database connection string.
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
		dbHost, dbUser, dbName, dbSSLMode, dbPassword)

	// Establish the database connection.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// If the database connection fails, panic with an error message.
		panic("failed to connect database")
	}

	return db
}
