package config

import (
	"MovieBack/internal/types"
	"context"
)

func ExecuteQuery(query string, params types.Slice) (any, error) {
	queryData, err := DB.Acquire(context.Background())
	if err != nil {
		return "", err
	}
	defer queryData.Release() // Ensure connection is released

	result, err := queryData.Exec(context.Background(), query, params...)
	if err != nil {
		return "", err
	}

	rows := result.RowsAffected()
	if rows >= 1 {
		return result, nil
	}

	return result, nil
}
