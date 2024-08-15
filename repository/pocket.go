package repository

import (
	"time"

	"github.com/AndresCRamos/go-worker-transaction/models"
)

var PocketData []models.Pocket

func GetPocketsForUser(userId int) []models.Pocket {
	time.Sleep(time.Second)
	userPockets := []models.Pocket{}
	for _, pocket := range PocketData {
		if pocket.UserID == userId {
			userPockets = append(userPockets, pocket)
		}
	}
	return userPockets
}
