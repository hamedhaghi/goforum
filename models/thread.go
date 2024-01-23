package models

import "gorm.io/gorm"

type Thread struct {
	gorm.Model
	Title       string
	Description string
	Posts       []Post
}
