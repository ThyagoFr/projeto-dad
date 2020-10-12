package main

import (
	"log"

	"ufc.com/dad/src/utils"

	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Println(err)
		log.Println("No .env file found")
	}

}

func main() {
	/*
		mux := controller.NewRouter()
		c := cors.New(cors.Options{
			AllowedMethods: []string{"POST", "GET", "DELETE", "PUT", "PATCH"},
		})
		handler := c.Handler(mux)
		log.Println("Server running on 8080 port ... ")
		migration.Migrate()
		log.Fatal(http.ListenAndServe(":8080", handler))
	*/
	// utils.UploadToS3("1380713", "8317317")
	message := utils.Message{
		Name:  "Thyago",
		To:    "thyagofr@alu.ufc.br",
		Token: "837ygbjdouhsdn",
	}
	err := utils.SendMessage(message)
	log.Println(err)
}
