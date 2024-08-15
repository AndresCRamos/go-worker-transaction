package mock

import (
	"github.com/AndresCRamos/go-worker-transaction/models"
	"github.com/AndresCRamos/go-worker-transaction/utils/random"
)

var counter int = 0

func ReturnMockPocket(userID int) models.Pocket {
	counter++
	return models.Pocket{
		ID:     counter,
		UserID: userID,
		Name:   random.RandSeq(10),
		Amount: random.RandInt(1000, -1000),
	}
}
