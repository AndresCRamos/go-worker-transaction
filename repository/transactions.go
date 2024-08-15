package repository

import (
	"time"

	"github.com/AndresCRamos/go-worker-transaction/models"
)

var TransactionData []models.Transaction

func GetTransactionsForUser(userID int) []models.Transaction {
	time.Sleep(time.Second)
	transactionsUser := []models.Transaction{}
	for _, trx := range TransactionData {
		if trx.UserID == userID {
			transactionsUser = append(transactionsUser, trx)
		}
	}
	return transactionsUser
}

func GetSortedTransactionsForUser(userID int) map[int][]models.Transaction {
	time.Sleep(time.Second)
	transactionsSortedUser := map[int][]models.Transaction{}
	transactionsUser := GetTransactionsForUser(userID)

	for _, trx := range transactionsUser {
		_, ok := transactionsSortedUser[trx.PocketID]
		if !ok {
			transactionsSortedUser[trx.PocketID] = []models.Transaction{}
		}
		transactionsSortedUser[trx.PocketID] = append(transactionsSortedUser[trx.PocketID], trx)
	}
	return transactionsSortedUser
}
