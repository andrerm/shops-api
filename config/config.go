package config

import (
	"ShopsAPI/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the DSN from environment variables
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set")
	}

	// Connect to the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Run AutoMigrate on the connected database
	err = DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Store{},
		&models.Product{},
		&models.Transaction{},
		&models.PaymentType{},
		&models.Wallet{},
		&models.UserRole{},
		&models.Bill{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
