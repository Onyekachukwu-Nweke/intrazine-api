package models

import (
	"time"
)

type UserRole string

const (
	RoleReader UserRole = "reader"
	RoleAuthor UserRole = "author"
	RoleEditor UserRole = "editor"
	RoleAdmin  UserRole = "admin"
)

type User struct {
	ID              string    `json:"id" db:"id"`
	Username        string    `json:"username" db:"username" validate:"required"`
	Email           string    `json:"email" db:"email" validate:"required"`
	Photo           string    `json:"photo" db:"photo"`
	PasswordHash    string    `json:"-" db:"password_hash"`
	Role            UserRole  `json:"role" db:"role"`
	TwoFactorSecret string    `json:"-" db:"two_factor_secret"`
	Is2FAEnabled    bool      `json:"is_2fa_enabled" db:"is_2fa_enabled"`
	LoginAttempts   int       `json:"-" db:"login_attempts"`
	LockedUntil     time.Time `json:"-" db:"locked_until"`
	CreatedAt       string    `json:"created_at" db:"created_at"`
	UpdatedAt       string    `json:"updated_at" db:"updated_at"`
}

// HasPermission checks if the user has sufficient permissions
func (u *User) HasPermission(requiredRole UserRole) bool {
	roleHierarchy := map[UserRole]int{
		RoleReader: 1,
		RoleAuthor: 2,
		RoleEditor: 3,
		RoleAdmin:  4,
	}

	return roleHierarchy[u.Role] >= roleHierarchy[requiredRole]
}
