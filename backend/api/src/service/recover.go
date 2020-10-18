package service

import (
	"errors"
	"math/rand"
	"time"

	"gorm.io/gorm"

	"ufc.com/dad/src/model"

	"ufc.com/dad/src/utils"
)

// SendEmail - SendEmail
func SendEmail(email string) error {

	result, err := getReaderData(email)
	if err != nil {
		return err
	}
	err = utils.SendMessage(result)
	if err != nil {
		return err
	}
	createRegister(result)
	return nil

}

// RecoverPassword - RecoverPassword
func RecoverPassword(token string, password string) error {

	reader, err := findRegister(token)
	if err != nil {
		return err
	}
	err = UpdatePassword(reader.Email, password)
	return err

}

// GetReaderData - GetReaderData
func getReaderData(email string) (utils.Message, error) {

	db, _ := utils.NewConnection()
	var data utils.Message
	err := db.Raw("SELECT name, email AS to FROM readers WHERE email = ?", email).Scan(&data)
	if err != nil {
		return data, errors.New("Usuario nao encontrado")
	}
	data.Token = generateRandomToken()
	return data, nil

}

// GenerateRandomToken - GenerateRandomToken
func generateRandomToken() string {

	const size = 8
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, size)
	for index := range s {
		s[index] = letters[rand.Intn(len(letters))]
	}
	return string(s)

}

// CreateRegister - RecoverPassword
func createRegister(data utils.Message) {

	db, _ := utils.NewConnection()
	var rec model.Recover
	err := db.First(&rec, "token = ? AND email = ? AND retrived IS NOT NULL", rec.Token, rec.Email).Error
	if err == gorm.ErrRecordNotFound {
		rec.Token = data.Token
		rec.Email = data.To
		db.Create(&rec)
	} else {
		rec.Retrieved = time.Now()
		db.Save(&rec)
	}
}

func findRegister(token string) (model.Recover, error) {

	db, _ := utils.NewConnection()
	var rec model.Recover
	err := db.First(&rec, "token = ? AND retrived IS NULL", rec.Token).Error
	if err == gorm.ErrRecordNotFound {
		return rec, err
	}
	return rec, nil

}
