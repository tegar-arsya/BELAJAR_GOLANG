// migrations/migrate.go

package migrations

import (
	"portfolio-backend/config"
	"portfolio-backend/models"
	"log"
)

func Migrate() {
	if err := config.DB.AutoMigrate(&models.TokenBlacklist{}); err != nil {
		log.Fatalf("Error migrating models: %v", err)
	}
}
