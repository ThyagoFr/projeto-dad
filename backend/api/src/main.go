package main

import (
	"log"
	"net/http"

	"ufc.com/dad/src/controller"
	"ufc.com/dad/src/migration"
	"ufc.com/dad/src/utils"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Println(err)
		log.Println("No .env file found")
	}

}

func main() {

	mux := controller.NewRouter()
	c := cors.AllowAll()

	handler := c.Handler(mux)
	log.Println("Server running on 8080 port ... ")
	migration.Migrate()
	utils.LoadInitalData()
	log.Fatal(http.ListenAndServe(":8080", handler))
}
