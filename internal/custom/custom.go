package custom

import "fmt"

func Map[T any](data []T) {

	for i, val := range data {
		fmt.Println(i, val)
	}

}
