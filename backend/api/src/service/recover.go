package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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
	err := db.Raw("SELECT name, email AS to FROM readers WHERE email = ?", email).Scan(&data).Error
	if err != nil {
		return data, errors.New("Usuario nao encontrado")
	}
	data.Token = generateRandomToken()
	return data, nil

}

// GenerateRandomToken - GenerateRandomToken
func generateRandomToken() string {

	const size = 12
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
	db.Raw("SELECT * FROM recovers WHERE email = ? AND retrieved IS NULL", data.To).Scan(&rec)
	if rec.ID != 0 {
		rec.Retrieved.Scan(time.Now())
		db.Save(&rec)
		log.Println(rec)
	}
	var rec2 model.Recover
	rec2.Token = data.Token
	rec2.Email = data.To
	rec2.Retrieved = sql.NullTime{}

	db.Create(&rec2)
	log.Println(rec2)

}

func findRegister(token string) (model.Recover, error) {

	db, _ := utils.NewConnection()
	var rec model.Recover
	fmt.Println(token)
	err := db.First(&rec, "token = ? AND retrieved IS NULL", token).Error
	if err == gorm.ErrRecordNotFound {
		return rec, errors.New("Token inválido")
	}
	return rec, nil

}
