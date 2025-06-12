package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/tensuqiuwulu/go-alokasir/cmd/initialize"
	"github.com/tensuqiuwulu/go-alokasir/cmd/variables"
	"github.com/tensuqiuwulu/go-alokasir/pkg/utilities"
)

func main() {
	// Load configuration
	viper.SetConfigFile(variables.CONFIG)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// Initialize logger
	isDev := viper.GetString("app.env") == "development"
	if err := utilities.InitializeLogger(isDev); err != nil {
		log.Fatalf("Failed to initialize logger: %s", err)
	}
	defer utilities.Logger.Sync()

	logger := utilities.GetLogger()
	logger.Info("Application starting...")

	// Initialize Echo
	e := echo.New()

	// Initialize database
	db := initialize.InitializeDatabase()

	// Initialize modules
	initialize.InitializeModules(e, db)

	// Display routes table
	go initialize.InitializeRoutes(e)

	// Start server
	serverAddr := fmt.Sprintf(":%d", viper.GetInt("app.port"))
	if err := e.Start(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}