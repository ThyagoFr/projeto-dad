package model

import (
	"gorm.io/gorm"
)

// Comment - Comment struct
type Comment struct {
	gorm.Model
	Comment  string
	Rate     float64
	BookID   uint
	ReaderID uint
}

// Comments - Comments
type Comments []Comment
