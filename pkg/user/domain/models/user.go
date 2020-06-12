package user

import "time"

// User model structure for User
type User struct {
	ID       int64
	Name     string
	Nickname string
	Phone    int
	CreateAt time.Time
	UpdateAt time.Time
}

// CreateUserCmd
type CreateUserCmd struct {
	Name     string    `json:"name"`
	Nickname string    `json:"nickname"`
	Phone    int       `json:"phone"`
}