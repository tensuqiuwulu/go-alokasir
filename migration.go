package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/tensuqiuwulu/go-alokasir/cmd/initialize"
	"github.com/tensuqiuwulu/go-alokasir/cmd/variables"
	"github.com/tensuqiuwulu/go-migration/migration"

	_ "plugin"
)

func main() {
		// Load configuration
		viper.SetConfigFile(variables.CONFIG)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}

		// Initialize database
		db := initialize.InitializeDatabase()

		// Initialize migrations
		if len(os.Args) > 1 && (os.Args[1] == "make:migration" || os.Args[1] == "migrate" || os.Args[1] == "migrate:rollback") {
			migration.SetDatabaseConnection(db)
	
			migration.ExecuteCommand(os.Args[1:])
			return
		}

		// need argument
		if len(os.Args) == 1 {
			fmt.Println("go-alokasir migration")
			fmt.Println("Available commands:")
			fmt.Println("  make:migration <name> - Create a new migration file")
			fmt.Println("  migrate - Run all pending migrations")
			fmt.Println("  migrate:rollback - Rollback the last batch of migrations")
			return
		}
}