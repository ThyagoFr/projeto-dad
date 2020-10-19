package service

import (
	"ufc.com/dad/src/model"
	"ufc.com/dad/src/utils"
)

// CommentResponse - CommentResponse
type CommentResponse struct {
	Email     string
	User      string
	UserPhoto string
	Comment   string
	Rate      float64
}

// CommentResponseList - CommentResponseList
type CommentResponseList []CommentResponse

// GetComments - Get all comments
func GetComments(idBook uint) CommentResponseList {

	var response CommentResponseList
	db, _ := utils.NewConnection()
	db.Raw(`SELECT 
			email,
			name as user,
			photo as user_photo,
			rate,
			comment
			FROM
			(SELECT * FROM readers) r
			INNER JOIN
			(SELECT * FROM comments WHERE book_id = ?) c
			ON c.reader_id = r.id
		  `, idBook).Scan(&response)
	return response

}

// StoreComment - Store a comment
func StoreComment(comment model.Comment) error {

	db, _ := utils.NewConnection()
	var bk model.Book
	db.Where("id = ?", comment.BookID).Find(&bk)
	err := db.Model(&bk).Association("Comments").Append(&comment)
	return err

}
