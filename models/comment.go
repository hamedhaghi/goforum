package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID  uint
	Content string
	Votes   int
}
