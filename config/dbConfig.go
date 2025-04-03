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

func IntializeDB() error {

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

	return err

}

func CloseDB() {

	if DB != nil {
		DB.Close()
		fmt.Println("database closed")
	}

}

// func ExecuteQuery(query string, params any) ([]map[string]interface{}, error) {
// 	if err != nil {
// 		return nil, fmt.Errorf("database is unreachable: %v", err)
// 	}
// 	client, err := db.Acquire(context.Background())
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to acquire connection: %v", err)
// 	}
// 	defer client.Release()
// 	rows, err := client.Query(context.Background(), query)
// 	if err != nil {
// 		return nil, fmt.Errorf("query execution failed: %v", err)
// 	}
// 	defer rows.Close()

// 	// Process the results into a slice of maps
// 	var results []map[string]interface{}
// 	for rows.Next() {
// 		values, err := rows.Values()
// 		if err != nil {
// 			return nil, err
// 		}

// 		rowMap := make(map[string]interface{})
// 		for i, field := range rows.FieldDescriptions() {
// 			rowMap[string(field.Name)] = values[i]
// 		}
// 		results = append(results, rowMap)
// 	}

// 	return results, nil
// }
