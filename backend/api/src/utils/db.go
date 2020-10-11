package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database - Database struct
type Database struct {
	dialect  string
	host     string
	port     string
	user     string
	dbname   string
	password string
}

// New -- Cria uma nova Struct com os valores de configuração
func (d *Database) New() {
	d.dialect = "postgres"
	d.host = "localhost"
	d.port = "5432"
	d.user = "appdad"
	d.password = "392035"
	d.dbname = "appdad"

}

// NewConnection -- Cria uma nova conexão com o banco
func NewConnection() (*gorm.DB, error) {
	databaseParams := Database{}
	databaseParams.New()
	stringConnection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		databaseParams.host, databaseParams.port,
		databaseParams.user, databaseParams.dbname,
		databaseParams.password)
	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  stringConnection,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
	return db, err
}
