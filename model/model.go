package model

import "time"

type User struct {
	UserId    string    `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	MobileNo  string    `json:"mobile_no"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}
