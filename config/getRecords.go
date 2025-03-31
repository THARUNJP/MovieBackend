package config

import (
	"MovieBack/internal/types"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

func GetRecords(query string, params types.Slice) ([]map[string]interface{}, error) {
	// Acquire connection from pool
	conn, err := DB.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to acquire connection: %v", err)
	}
	defer conn.Release()

	// Execute query
	rows, err := conn.Query(context.Background(), query, params...)
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	// Prepare result slice
	var results []map[string]interface{}
	fieldDescriptions := rows.FieldDescriptions()

	// Process each row
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("failed to get row values: %v", err)
		}

		row := make(map[string]interface{})
		for i, fd := range fieldDescriptions {
			colName := string(fd.Name)
			val := values[i]

			// Convert UUID bytes to string if needed
			if fd.DataTypeOID == pgtype.UUIDOID {
				if uuidBytes, ok := val.([16]byte); ok {
					row[colName] = fmt.Sprintf("%x-%x-%x-%x-%x",
						uuidBytes[0:4],
						uuidBytes[4:6],
						uuidBytes[6:8],
						uuidBytes[8:10],
						uuidBytes[10:16])
					continue
				}
			}
			row[colName] = val
		}
		results = append(results, row)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return results, nil
}
