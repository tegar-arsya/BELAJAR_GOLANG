// migrations/migrate.go
package migrations

import (
    "portfolio-backend/config"
    "portfolio-backend/models"
)

func Migrate() {
    config.DB.AutoMigrate( &models.Portfolio{}, &models.About{}, &models.User{}, &models.Cv{}, &models.Pengalaman{}, &models.Sertifikat{}, &models.Article{} )
}
