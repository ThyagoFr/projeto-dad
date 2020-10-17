package service

import (
	"log"

	"ufc.com/dad/src/model"
	"ufc.com/dad/src/utils"
)

// CommentResponse - CommentResponse
type CommentResponse struct {
	Author      string
	AuthorPhoto string
	Comment     string
	Rate        float64
}

// CommentResponseList - CommentResponseList
type CommentResponseList []CommentResponse

// GetComments - Get all comments
func GetComments(idBook uint) CommentResponseList {

	var comments model.Comments
	db, _ := utils.NewConnection()
	err := db.Model(&model.Comment{}).Where("book_id = ?", idBook).Find(&comments).Error
	if err != nil {
		log.Fatal(err)
	}
	return AllCommentsToResponseList(comments)

}

// getAVGRateAndNumberOfComments - getAVGRateAndNumberOfComments
func getAVGRateAndNumberOfComments(idBook uint) (int, float64) {

	type Result struct {
		Number int
		Sum    float64
	}
	var result Result
	avg := 0.0
	db, _ := utils.NewConnection()
	db.Raw("SELECT SUM(rate) as sum, COUNT(*) AS number FROM comments WHERE book_id = ?", idBook).Scan(&result)
	if result.Number != 0 {
		avg = result.Sum / float64(result.Number)
	}
	return result.Number, avg
}

// StoreComment - Store a comment
func StoreComment(comment model.Comment) error {

	db, _ := utils.NewConnection()
	db.Create(&comment)
	return nil

}

// AllCommentsToResponseList - AllCommentsToResponseList
func AllCommentsToResponseList(comments model.Comments) CommentResponseList {

	var response CommentResponseList
	for _, com := range comments {
		reader, _ := GetOneReader(com.ReaderID)
		res := CommentResponse{
			Comment:     com.Comment,
			Rate:        com.Rate,
			Author:      reader.Name,
			AuthorPhoto: reader.Photo,
		}
		response = append(response, res)
	}
	return response

}
