package config

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("DB_CONNECTION_STRING environment variable is not set")
	}

	m, err := migrate.New(
		"file://migration",
		connectionString,
	)
	if err != nil {
		log.Fatalf("failed to initialize migration: %v", err.Error())
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migration: %v", err)
	}

	log.Println("Database migration completed successfully")
}
