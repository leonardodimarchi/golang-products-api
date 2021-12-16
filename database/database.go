package database

import (
	"golang-products-api/database/migrations"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func StartDatabase() {
	database = createDatabaseInstance()
	configureDatabase()
}

func createDatabaseInstance() *gorm.DB {
	godotenv.Load(".env")

	databaseHost := os.Getenv("DB_HOST")
	databasePort := os.Getenv("DB_PORT")
	databaseUser := os.Getenv("DB_USER")
	databasePassword := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_DATABASE")

	databaseConnectionUrl := 
	"host=" + databaseHost + 
	" port=" + databasePort +
	" user=" + databaseUser + 
	" dbname=" + databaseName +
	" password=" + databasePassword +
	" sslmode=disable"

	createdDatabase, errorFromDatabaseOpen := gorm.Open(postgres.Open(databaseConnectionUrl), &gorm.Config{})

	if errorFromDatabaseOpen != nil {
		log.Fatal("Failed to create the database: ", errorFromDatabaseOpen)
	}

	return createdDatabase
}

func configureDatabase() {
	config, errorFromConfig := database.DB()

	if errorFromConfig != nil {
		log.Fatal("Failed to configure the database: ", errorFromConfig)
	}

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(database)
}

func GetDatabase() *gorm.DB {
	return database
}