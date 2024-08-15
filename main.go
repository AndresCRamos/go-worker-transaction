package main

import (
	"fmt"

	"github.com/AndresCRamos/go-worker-transaction/mock"
	"github.com/AndresCRamos/go-worker-transaction/models"
	"github.com/AndresCRamos/go-worker-transaction/pkg/worker_pool"
	"github.com/AndresCRamos/go-worker-transaction/repository"
)

type Response struct {
	UserData    models.User
	UserPockets []models.Pocket
	UserTrx     map[int][]models.Transaction
}

func main() {
	// job := func(number int) int {
	// 	min := 1000
	// 	max := min * 5
	// 	sleepTime := rand.Intn(max-min+1) + min
	// 	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	// 	return number * number
	// }

	mock.GenerateData()
	job := func(id int) Response {
		user, _ := repository.GetUserData(id)
		return Response{
			UserData:    user,
			UserPockets: repository.GetPocketsForUser(id),
			UserTrx:     repository.GetSortedTransactionsForUser(id),
		}
	}

	workerPool := worker_pool.NewWorkerPool(job, 12)
	inputCh := workerPool.GetInputChannel()
	resChan := workerPool.Run()

	// // Feed data into the pool
	// go func() {
	// 	for i := 1; i <= 100; i++ {
	// 		inputCh <- i
	// 	}
	// 	close(inputCh)
	// }()

	// Feed data into the pool
	go func() {
		for _, user := range repository.UserData {
			inputCh <- user.ID
		}
		close(inputCh)
	}()

	for res := range resChan {
		fmt.Println(res.UserData)
	}
}
