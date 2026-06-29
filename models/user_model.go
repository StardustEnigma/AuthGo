package models

import "time"

type User struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}