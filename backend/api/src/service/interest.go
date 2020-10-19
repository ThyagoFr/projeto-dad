package service

import (
	"ufc.com/dad/src/model"
	"ufc.com/dad/src/utils"
)

// GetInterests - GetInterests
func GetInterests(idReader uint) []BookInfo {

	result := []BookInfo{}
	db, _ := utils.NewConnection()
	db.Raw(`select id, interestid ,  title, cover, genre,	author , summary, average, qtd 
			from
			(
			(select book_id, id as interestid from interests where reader_id = ? AND deleted_at IS NULL) i
			inner join
			(select * from books) b
			on b.id = i.book_id
			) res
			left join 
			( select book_id, sum(rate)/count(book_id) as average, count(book_id) as qtd from comments c group by book_id) c
			on c.book_id = res.book_id`, idReader).Scan(&result)

	return result

}

// StoreInterest - StoreInterest
func StoreInterest(interest model.Interest) error {

	db, _ := utils.NewConnection()
	var reader model.Reader
	db.Where("id = ?", interest.ReaderID).Find(&reader)
	err := db.Model(&reader).Association("Interests").Append(&interest)
	return err

}

// RemoveInterest - RemoveInterest
func RemoveInterest(idU, idI uint) error {

	db, _ := utils.NewConnection()
	var interest model.Interest
	db.Where("id = ?", idI).Find(&interest)
	err := db.Delete(interest).Error
	return err

}
