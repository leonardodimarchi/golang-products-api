package migrations

import (
	"golang-products-api/models"

	"gorm.io/gorm"
)

func RunMigrations(database *gorm.DB) {
	database.AutoMigrate(models.Product{})
}
