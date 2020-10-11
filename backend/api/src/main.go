package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/rs/cors"
	"ufc.com/dad/src/migration"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

}

func main() {
	mux := NewRouter()
	c := cors.New(cors.Options{
		AllowedMethods: []string{"POST", "GET", "DELETE", "PUT", "PATCH"},
	})
	handler := c.Handler(mux)
	fmt.Println("Server running on 8080 port ... ")
	migration.Migrate()
	log.Fatal(http.ListenAndServe(":8080", handler))
}
