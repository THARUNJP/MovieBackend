package types

import (
	"time"

	"github.com/google/uuid"
)

type Slice []interface{}

type UserStruct struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	MobileNo string `json:"mobileno"`
	Password string `json:"password"`
}

type MovieStruct struct {
	MovieID    uuid.UUID `json:"movie_id"`
	MovieName  string    `json:"movie_name"`
	MovieGenre string    `json:"movie_genre"`
	ImgUrl     string    `json:"img_url"`
	IsActive   bool      `json:"is_active"`
}

type MovieIdRequest struct {
	ID string `json:"id"`
}

type UserRefrestToken struct {
	RefToken  string        `json:"value"`
	ExpiresAt time.Duration `json:"ttl"`
}
