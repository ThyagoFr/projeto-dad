package service

import (
	"log"

	"ufc.com/dad/src/model"
	"ufc.com/dad/src/utils"
)

// Comments - Comments
type Comments []model.Comment

// GetAllComments - Get all comments
func GetAllComments(idBook uint) Comments {

	var comments Comments
	db, _ := utils.NewConnection()
	err := db.Where("book_id = ?", idBook).Find(&comments).Error
	if err != nil {
		log.Fatal(err)
	}
	return comments

}

// StoreComment - Store a comment
func StoreComment(idBook uint, comment model.Comment) error {

	_, err := GetOneBook(idBook)
	if err != nil {
		return err
	}
	db, _ := utils.NewConnection()
	db.Create(&comment)
	return nil

}
