package service

import (
	"errors"

	"ufc.com/dad/src/model"
	"ufc.com/dad/src/security"
	"ufc.com/dad/src/utils"
)

// Readers - Readers
type Readers []model.Reader

// Login - Login
func Login(email, pass string) (string, error) {

	db, _ := utils.NewConnection()
	var reader model.Reader
	err := db.Where("email = ?", email).Find(&reader).Error
	if err != nil {
		return "", err
	}
	match := security.CheckPasswordHash(pass, reader.Password)
	if match {
		return security.GenerateToken(reader.ID)
	}
	return "", errors.New("Check your credentials")

}

// GetAllReaders - Get all readers
func GetAllReaders() Readers {

	var readers Readers
	db, _ := utils.NewConnection()
	db.Find(&readers)
	return readers

}

// GetOneReader - Get one specific reader
func GetOneReader(id uint) (*model.Reader, error) {

	db, _ := utils.NewConnection()
	var reader model.Reader
	err := db.Find(&reader).Error
	if err != nil {
		return nil, err
	}
	return &reader, nil

}

// StoreReader - Store a reader
func StoreReader(reader model.Reader) *model.Reader {

	db, _ := utils.NewConnection()
	reader.Password, _ = security.HashPassword(reader.Password)
	db.Create(&reader)
	return &reader

}

// DeleteReader - Delete a reader
func DeleteReader(id int) error {

	db, _ := utils.NewConnection()
	var reader model.Reader
	err := db.Where("id = ?", id).Find(&reader).Error
	db.Delete(&reader)
	if err != nil {

	}
	return nil
}

// UpdatePassword - UpdatePassword
func UpdatePassword(email, password string) error {

	db, _ := utils.NewConnection()
	reader, err := findReaderByEmail(email)
	if err != nil {
		return err
	}
	reader.Password, _ = security.HashPassword(password)
	db.Save(&reader)
	return nil

}

func findReaderByEmail(email string) (model.Reader, error) {

	db, _ := utils.NewConnection()
	var reader model.Reader
	err := db.Where("email = ?", email).Find(&reader).Error
	if err != nil {
		return reader, errors.New("Usuario nao encontrado")
	}
	return reader, nil

}
