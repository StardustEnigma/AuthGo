package models

import "time"

type UserRole string
const (
	Admin UserRole="admin"
	Users UserRole="user"
	Manager UserRole="manager"
)

type User struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	Role         UserRole    `json:"role"`     
    IsActive     bool      `json:"is_active"`
    IsVerified   bool      `json:"is_verified"`

    LastLoginAt  time.Time `json:"last_login_at"`
    UpdatedAt    time.Time `json:"updated_at"`
    DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}