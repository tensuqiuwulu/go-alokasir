package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tensuqiuwulu/go-alokasir/pkg/utilities"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() *gorm.DB {
	logger := utilities.GetLogger()

	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s timezone=%s",
		viper.GetString("pgsql.host"),
		viper.GetString("pgsql.username"),
		viper.GetString("pgsql.password"),
		viper.GetString("pgsql.dbname"),
		viper.GetInt("pgsql.port"),
		viper.GetString("pgsql.sslmode"),
		viper.GetString("pgsql.timezoneconf.timezone"),
	)

	// Open connection to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to database",
			zap.Error(err),
			zap.String("host", viper.GetString("pgsql.host")),
			zap.Int("port", viper.GetInt("pgsql.port")),
		)
	}

	logger.Info("Database connection established successfully")

	// Get the underlying *sql.DB instance
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("Failed to get database instance",
			zap.Error(err),
		)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(viper.GetInt("pgsql.max_idle_connections"))
	sqlDB.SetMaxOpenConns(viper.GetInt("pgsql.max_open_connections"))

	// Enable UUID extension
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	return db
}