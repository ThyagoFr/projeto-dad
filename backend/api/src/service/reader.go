package service

import (
	"errors"
	"mime/multipart"

	"ufc.com/dad/src/model"
	"ufc.com/dad/src/security"
	"ufc.com/dad/src/utils"
)

// Readers - Readers
type Readers []model.Reader

// Login - Login
func Login(email, pass string) (string, *model.Reader, error) {

	db, _ := utils.NewConnection()
	var reader model.Reader
	err := db.Where("email = ?", email).Find(&reader).Error
	if err != nil {
		return "", nil, err
	}
	match := security.CheckPasswordHash(pass, reader.Password)
	if match {
		token, error := security.GenerateToken(reader.ID)
		return token, &reader, error
	}
	return "", nil, errors.New("Check your credentials")

}

// GetAllReaders - Get all readers
func GetAllReaders() Readers {

	var readers Readers
	db, _ := utils.NewConnection()
	db.Select("id", "name", "email", "photo", "age").Find(&readers)
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
func StoreReader(reader model.Reader, file multipart.File) (*model.Reader, error) {

	db, _ := utils.NewConnection()
	_, err := findReaderByEmail(reader.Email)
	if err == nil {
		return &reader, errors.New("Email ja utilizado")
	}
	reader.Password, _ = security.HashPassword(reader.Password)
	db.Create(&reader)
	reader.Photo, _ = utils.UploadReaderProfileToS3(reader.ID, file)
	db.Save(&reader)
	return &reader, nil

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

// UpdateReader - UpdateReader
func UpdateReader(reader model.Reader) error {

	db, _ := utils.NewConnection()
	err := db.Where("id = ?", reader.ID).Find(&reader).Error
	if err != nil {
		return errors.New("Usuario nao encontrado")
	}
	db.Model(&reader).Updates(&reader)
	return nil

}

func findReaderByEmail(email string) (model.Reader, error) {

	db, _ := utils.NewConnection()
	var reader model.Reader
	db.Where("email = ?", email).Find(&reader)
	if reader.Email == "" {
		return reader, errors.New("Usuario nao encontrado")
	}
	return reader, nil

}
