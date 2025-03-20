package types

import "encoding/json"

type Slice []interface{}

type UserStruct struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	MobileNo string `json:"mobileno"`
	Password string `json:"password"`
}

type MovieStruct struct {
	MovieName    string          `json:"movie_name"`
	MovieDetails json.RawMessage `json:"movie_details"` // Handles JSON data
}
