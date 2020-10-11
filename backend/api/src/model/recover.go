package model

import (
	"time"

	"gorm.io/gorm"
)

// Recover - Recover struct
type Recover struct {
	gorm.Model
	Email     string
	Token     string
	Retrieved time.Time
}
