package config

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func DownMigrations() {
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("DB_CONNECTION_STRING environment variable is not set")
	}

	m, err := migrate.New(
		"file://../migration", // Adjust path based on where you run the app
		connectionString,
	)
	if err != nil {
		log.Fatalf("failed to initialize migration: %v", err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migration: %v", err.Error())
	}

	log.Println("Database migration completed successfully")
}
