package initialize

import (
	"github.com/tensuqiuwulu/go-migration/migration"
	"gorm.io/gorm"
)

func InitializeMigrations(db *gorm.DB) {
	migration.SetDatabaseConnection(db)
}
