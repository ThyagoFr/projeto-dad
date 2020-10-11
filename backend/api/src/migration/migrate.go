package migration

import (
	"log"

	"ufc.com/dad/src/config"
	"ufc.com/dad/src/model"
)

// Migrate - Migrate function
func Migrate() {

	db, err := config.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(
		&model.Reader{},
		&model.Recover{},
		&model.Book{},
		&model.Comment{},
		&model.Interest{},
	)

}
