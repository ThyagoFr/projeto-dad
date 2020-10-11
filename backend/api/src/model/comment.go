package model

import (
	"gorm.io/gorm"
)

// Comment - Comment struct
type Comment struct {
	gorm.Model
	Comment  string
	Rate     uint
	BookID   uint
	ReaderID uint
}
