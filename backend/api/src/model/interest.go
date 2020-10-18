package model

import (
	"gorm.io/gorm"
)

// Interest - Interest struct
type Interest struct {
	gorm.Model
	BookID   uint `json:"book_id"`
	ReaderID uint
}

// Interests - Interests
type Interests []Interest
