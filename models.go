package main

import (
	"time"
)

// Bookmark model used by database
type Bookmark struct {
	ID          uint   `json:"id" gorm:"primary_key" sql:"AUTO_INCREMENT"`
	Name        string `json:"name" sql:"type:varchar(36)"`
	Description string `json:"description" sql:"type:varchar(200)"`
	URL         string `json:"url" sql:"type:varchar(300)"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// User model used by database
type User struct {
	ID           uint   `json:"id" gorm:"primary_key" sql:"AUTO_INCREMENT"`
	Username     string `json:"username" gorm:"primary_key" sql:"type:varchar(30);unique"`
	FirstName    string `json:"first_name" sql:"type:varchar(30)"`
	LastName     string `json:"last_name" sql:"type:varchar(30)"`
	Email        string `json:"email" sql:"type:varchar(30)"`
	PasswordHash string `json:"password_hash" sql:"type:varchar(60)"`
	Role         string `json:"role" sql:"type:varchar(20)"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
