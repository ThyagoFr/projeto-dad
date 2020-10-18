package model

import (
	"gorm.io/gorm"
)

// Comment - Comment struct
type Comment struct {
	gorm.Model
	Comment  string  `json:"comment"`
	Rate     float64 `json:"rate"`
	BookID   uint    `json:"book_id"`
	ReaderID uint    `json:"reader_id"`
}

// Comments - Comments
type Comments []Comment
