package config

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DB               *pgxpool.Pool
	once             sync.Once
	connectionString string
)

func IntializeDB() (*pgxpool.Pool, error) {

	connectionString = os.Getenv("DB_CONNECTION_STRING")

	var err error
	once.Do(func() {
		DB, err = pgxpool.New(context.Background(), connectionString)

		if err != nil {
			fmt.Printf("error connecting to db %v", err)
		}
		dbError := DB.Ping(context.Background())
		if dbError != nil {
			fmt.Printf("db error %v", dbError)
		}

		fmt.Println("db connected successfully")
	})

	return DB, err

}

func CloseDB() {

	if DB != nil {
		DB.Close()
		fmt.Println("database closed")
	}

}
