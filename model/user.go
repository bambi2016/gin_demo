package model

import "time"

// User represents a user entity
type User struct {
	ID        uint64     `json:"id"`
	Username  string     `json:"username" binding:"required,min=3,max=20"`
	Email     string     `json:"email" binding:"required,email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
