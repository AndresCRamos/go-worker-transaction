package mock

import (
	"math/rand"

	"github.com/AndresCRamos/go-worker-transaction/models"
	"github.com/AndresCRamos/go-worker-transaction/utils/random"
)

var counterTrx int = 0

func ReturnMockTransaction(userID int, pocketID int) models.Transaction {
	counterTrx++
	return models.Transaction{
		ID:       counterTrx,
		Name:     random.RandSeq(10),
		Date:     random.RandDateLast30Days(),
		Type:     uint(2),
		Amount:   uint(rand.Intn(1000)),
		UserID:   userID,
		PocketID: pocketID,
	}
}
