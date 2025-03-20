package config

import (
	"context"
	"fmt"
)

func GetRecords(query string) (any, error) {

	rows, err := DB.Acquire(context.Background())

	if err != nil {
		return fmt.Printf("err in connecting")
	}
	defer rows.Release()
	result, err := rows.Query(context.Background(), query)

	if err != nil {
		return fmt.Printf("err in query %v", err)

	}

	var data []map[string]any // Slice to store all rows

	columns := result.FieldDescriptions()
	for result.Next() {
		values, err := result.Values()
		if err != nil {
			return nil, fmt.Errorf("error in iteration: %v", err)
		}

		rowData := make(map[string]any) // Map to store row data

		for i := range columns {
			rowData[string(columns[i].Name)] = values[i] // Convert column name to string
		}

		data = append(data, rowData) // Append row data to the slice
	}

	errCheck := result.Err()
	if errCheck != nil {
		return nil, fmt.Errorf("erroor in query %v", errCheck)
	}

	return data, nil

}
