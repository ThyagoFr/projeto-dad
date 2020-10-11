package service

import (
	"ufc.com/dad/src/config"
	"ufc.com/dad/src/model"
)

// Readers - Readers
type Readers []model.Reader

// GetAllReaders - Get all readers
func GetAllReaders() Readers {

	var readers Readers
	db, _ := config.NewConnection()
	db.Find(&readers)

	return readers
}

// GetOneReader - Get one specific reader
func GetOneReader(id int) (*model.Reader, error) {

	db, _ := config.NewConnection()
	var reader model.Reader
	err := db.Where("id = ?", id).Find(&reader).Error
	if err != nil {
		return nil, err
	}
	return &reader, nil

}

// StoreReader - Store a reader
func StoreReader(reader model.Reader) *model.Reader {

	db, _ := config.NewConnection()
	reader.Password, _ = config.HashPassword(reader.Password)
	db.Create(&reader)
	return &reader

}

// DeleteReader - Delete a reader
func DeleteReader(id int) error {

	db, _ := config.NewConnection()
	var reader model.Reader
	err := db.Where("id = ?", id).Find(&reader).Error
	db.Delete(&reader)
	if err != nil {

	}
	return nil
}
