package models

import "time"

type UserStatus string

const (
	AdminUser   UserStatus = "admin"
	DefaultUser UserStatus = "default"
)

type User struct {
	ID          int64
	Email       string
	FirstName   string
	LastName    string
	Surname     string
	AvatarURL   string
	Status      UserStatus
	CreatedAt   time.Time
	DateOfBirth time.Time
}
