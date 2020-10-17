package model

import (
	"gorm.io/gorm"
)

// Reader - Reader struct
type Reader struct {
	gorm.Model
	Name      string
	Email     string
	Age       uint
	Photo     string
	Password  string
	Interests []Interest
}

// Readers - Readers
type Readers []Reader
