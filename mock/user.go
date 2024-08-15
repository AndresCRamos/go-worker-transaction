package mock

import (
	"github.com/AndresCRamos/go-worker-transaction/models"
	"github.com/AndresCRamos/go-worker-transaction/utils/random"
)

func ReturnMockUser(id int) models.User {
	return models.User{
		ID:        id,
		Username:  random.RandSeq(10),
		LoginDate: random.RandDateLast30Days(),
	}
}
