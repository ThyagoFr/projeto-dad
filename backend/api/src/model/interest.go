package model

import (
	"gorm.io/gorm"
)

// Interest - Interest struct
type Interest struct {
	gorm.Model
	BookID   uint
	ReaderID uint
}

// Interests - Interests
type Interests []Interest
