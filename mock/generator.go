package mock

import (
	"math/rand"

	"github.com/AndresCRamos/go-worker-transaction/models"
	"github.com/AndresCRamos/go-worker-transaction/repository"
	"github.com/AndresCRamos/go-worker-transaction/utils/random"
)

func GenerateUser() (models.User, []models.Pocket, []models.Transaction) {
	user := ReturnMockUser(rand.Intn(100))
	pocketList := []models.Pocket{}
	trxList := []models.Transaction{}

	for range 10 {
		pocket := ReturnMockPocket(user.ID)

		for range random.RandInt(10, 1) {
			trx := ReturnMockTransaction(user.ID, pocket.ID)
			if trx.Type == models.TransactionCost {
				pocket.Amount -= int(trx.Amount)
			} else {
				pocket.Amount += int(trx.Amount)
			}
			trxList = append(trxList, trx)
		}

		pocketList = append(pocketList, pocket)
	}

	return user, pocketList, trxList
}

func GenerateData() {
	for range 100 {
		user, pocketList, trxList := GenerateUser()
		repository.UserData = append(repository.UserData, user)
		repository.PocketData = append(repository.PocketData, pocketList...)
		repository.TransactionData = append(repository.TransactionData, trxList...)
	}
}
