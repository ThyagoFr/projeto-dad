package service

import (
	"errors"
	"math/rand"

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
	return nil

}

// GetReaderData - GetReaderData
func getReaderData(email string) (utils.Message, error) {

	db, _ := utils.NewConnection()
	var reader model.Reader
	var message utils.Message
	err := db.Raw("SELECT * FROM readers WHERE email = ?", email).Scan(&reader).Error
	if err != nil {
		return message, errors.New("Usuario nao encontrado")
	}
	message.Token = generateRandomToken()
	message.To = reader.Email
	message.Name = reader.Name
	UpdatePassword(reader.Email, message.Token)
	return message, nil

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
