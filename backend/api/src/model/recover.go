package model

import (
	"database/sql"

	"gorm.io/gorm"
)

// Recover - Recover struct
type Recover struct {
	gorm.Model
	Email     string
	Token     string
	Retrieved sql.NullTime
}
