package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	// Get the database URL from environment variable
	dsn := os.Getenv("DB")
	if dsn == "" {
		log.Fatal("Environment variable DB is not set")
	}

	// Connect to the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database connection established")
}
