package main

// Bookmark model used by database
type Bookmark struct {
	ID          uint   `json:"id" gorm:"primary_key" sql:"AUTO_INCREMENT"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
}