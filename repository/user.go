package repository

import (
	"errors"
	"time"

	"github.com/AndresCRamos/go-worker-transaction/models"
)

var UserData []models.User

func GetUserData(userID int) (models.User, error) {
	time.Sleep(time.Second)
	for _, user := range UserData {
		if user.ID == userID {
			return user, nil
		}
	}
	return models.User{}, errors.New("Cant find user")
}
