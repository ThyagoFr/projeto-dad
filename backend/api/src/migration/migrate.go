package migration

import (
	"log"

	"ufc.com/dad/src/model"
	"ufc.com/dad/src/utils"
)

// Migrate - Migrate function
func Migrate() {

	db, err := utils.NewConnection()
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
