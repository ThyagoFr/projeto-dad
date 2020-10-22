package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

// Database - Database struct
type database struct {
	dialect  string
	host     string
	port     string
	user     string
	dbname   string
	password string
}

func init() {

	databaseParams := database{}
	databaseParams.new()
	stringConnection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		databaseParams.host, databaseParams.port,
		databaseParams.user, databaseParams.dbname,
		databaseParams.password)
	connection, _ = gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  stringConnection,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})

}

// NewConnection -- Cria uma nova conexão com o banco
func NewConnection() (*gorm.DB, error) {

	return connection, nil

}

// New -- Cria uma nova Struct com os valores de configuração
func (d *database) new() {
	d.dialect = "postgres"
	d.host = "appdad.chct2onrz0xz.us-east-1.rds.amazonaws.com"
	d.port = "5432"
	d.user = "postgres"
	d.password = "392035ab"
	d.dbname = "appdad"

}
