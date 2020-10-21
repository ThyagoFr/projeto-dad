package main

import (
	"log"
	"net/http"

	"ufc.com/dad/src/controller"
	"ufc.com/dad/src/migration"

	"github.com/rs/cors"
)

func main() {

	mux := controller.NewRouter()
	c := cors.AllowAll()

	handler := c.Handler(mux)
	log.Println("Server running on 8080 port ... ")
	migration.Migrate()
	// utils.LoadInitalData()
	log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", handler))
}
