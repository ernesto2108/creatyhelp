package domain

import "time"

// User model structure for User
type User struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Nickname string    `json:"nickname"`
	Phone    int       `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserCmd
type CreateUserCmd struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Phone    int    `json:"phone"`
}

// UpdateUserCmd
type UpdateUserCmd struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Phone    int    `json:"phone"`
}
