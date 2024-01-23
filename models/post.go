package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ThreadID uint
	Title    string
	Content  string
	Votes    int
	Comments []Comment
}
