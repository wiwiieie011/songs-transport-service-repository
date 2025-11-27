package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDatabaseConnection() *gorm.DB {
    // Загружаем .env
    if err := godotenv.Load(".env"); err != nil {
        panic(fmt.Sprintf("Error loading .env file: %v", err))
    }

    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASSWORD") // ⚡ нужно совпадать с .env
    dbHost := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")
    dbSSLMode := os.Getenv("DB_SSLMODE")
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        dbHost, dbUser, dbPass, dbName, dbPort, dbSSLMode,
    )

    db, err := gorm.Open(postgres.New(postgres.Config{
        DSN:                  dsn,
        PreferSimpleProtocol: true,
    }), &gorm.Config{})

    if err != nil {
        panic(fmt.Sprintf("failed to connect database: %v", err))
    }

    return db
}