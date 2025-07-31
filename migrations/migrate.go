// migrations/migrate.go
package migrations

import (
    "portfolio-backend/config"
    "portfolio-backend/internal/domain/entity"
)

func Migrate() {
    config.DB.AutoMigrate( &entity.Portfolio{}, &entity.About{}, &entity.User{}, &entity.Cv{}, &entity.Pengalaman{}, &entity.Sertifikat{}, &entity.Article{} )
}
