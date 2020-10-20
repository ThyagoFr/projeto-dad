package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"ufc.com/dad/src/model"
)

// Response - Response
type Response struct {
	Kind       string
	TotalItems int
	Items      []struct {
		VolumeInfo struct {
			Title       string
			Authors     []string
			Categories  []string
			Description string
			ImageLinks  struct {
				SmallThumbnail string
			}
			IndustryIdentifiers []struct {
				Identifier string
			}
		}
	}
}

// LoadInitalData - Load the initial book data to the database
func LoadInitalData() {

	db, _ := NewConnection()

	subjects := []string{
		"fiction",
		"horror",
		"romance",
		"adventure",
		"drama",
	}

	for _, subject := range subjects {

		url := "https://www.googleapis.com/books/v1/volumes?q=subject:" + subject + "&langRestrict=pt&maxResults=20"
		response, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		defer response.Body.Close()
		data := Response{}
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalln(err)
		}
		json.Unmarshal(body, &data)

		for _, element := range data.Items {
			title := element.VolumeInfo.Title
			cover := element.VolumeInfo.ImageLinks.SmallThumbnail

			if !strings.Contains(cover, "http://") {
				break
			}

			genre := "Desconhecido"
			author := "Desconhecido"

			if len(element.VolumeInfo.Authors) > 0 {
				author = element.VolumeInfo.Authors[0]
			}
			if len(element.VolumeInfo.Categories) > 0 {
				genre = element.VolumeInfo.Categories[0]
			}

			summary := element.VolumeInfo.Description
			url := ""
			if len(element.VolumeInfo.IndustryIdentifiers) != 0 {
				url, _ = UploadBookCoverToS3(element.VolumeInfo.IndustryIdentifiers[0].Identifier, cover)
			}
			book := model.Book{
				Title:   title,
				Cover:   url,
				Genre:   genre,
				Author:  author,
				Summary: summary,
			}

			os.Remove("temp.jpeg")
			db.Create(&book)
		}

	}
	log.Println("Carga inicial completa...")

}
