package service

import (
	"ufc.com/dad/src/model"
	"ufc.com/dad/src/utils"
)

// BookInfo - BookInfo
type BookInfo struct {
	ID      int
	Title   string
	Cover   string
	Genre   string
	Author  string
	Summary string
	Average float64
	Qtd     int
}

// GetBooks - Get all books
func GetBooks() []BookInfo {

	var result []BookInfo
	db, _ := utils.NewConnection()
	db.Raw(`select id, title, cover, genre,	author , summary, average, qtd 
			from
			books b
			left join
			( select book_id, sum(rate)/count(book_id) as average, count(book_id) as qtd from comments c group by book_id) c
			on c.book_id = b.id`).Scan(&result)
	return result

}

// GetBook - Get one specific book
func GetBook(id uint) (*model.Book, error) {

	db, _ := utils.NewConnection()
	var book model.Book
	err := db.Where("id = ?", id).Find(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil

}
