package model

import (
	"gorm.io/gorm"
)

// Book - Book struct
type Book struct {
	gorm.Model
	Title    string
	Cover    string
	Genre    string
	Author   string
	Summary  string
	Comments []Comment
}
