package service

import (
	"ufc.com/dad/src/model"
	"ufc.com/dad/src/utils"
)

// Books - Books
type Books []model.Book

// GetAllBooks - Get all books
func GetAllBooks() Books {

	var books Books
	db, _ := utils.NewConnection()
	db.Find(&books)
	return books

}

// GetOneBook - Get one specific book
func GetOneBook(id uint) (*model.Book, error) {

	db, _ := utils.NewConnection()
	var book model.Book
	err := db.Where("id = ?", id).Find(&book).Error
	if err != nil {
		return nil, err
	}
	book.Comments = GetAllComments(book.ID)
	return &book, nil

}
