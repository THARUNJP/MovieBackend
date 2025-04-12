package config

import (
	"MovieBack/internal/types"
	"context"
	"fmt"
)

func GetUsers(query string, params types.Slice) ([]types.MovieStruct, error) {

	conn, err := DB.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()

	// Execute the query with provided parameters
	rows, err := conn.Query(context.Background(), query, params...)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var users []types.MovieStruct
	for rows.Next() {
		var u types.MovieStruct
		err = rows.Scan(&u.MovieID, &u.MovieName, &u.MovieGenre, &u.IsActive)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil
}
