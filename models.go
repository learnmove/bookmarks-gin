package main

import (
	"time"
)

// Bookmark model used by database
type Bookmark struct {
	ID          uint   `json:"id" gorm:"primary_key" sql:"AUTO_INCREMENT"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`

	CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt *time.Time `json:"deleted_at"`
}

// User model used by database
type User struct {
	ID           uint   `json:"id" gorm:"primary_key" sql:"AUTO_INCREMENT"`
	Username     string `json:"username" gorm:"primary_key" sql:"unique"`
	FirstName     string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role     string `json:"role"`

	CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt *time.Time `json:"deleted_at"`
}
