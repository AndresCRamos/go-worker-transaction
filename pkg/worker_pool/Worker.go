package worker_pool

import (
	"fmt"
	"sync"
)

type worker[I, O any] struct {
	id     int
	input  chan I
	output chan O
	job    func(I) O
	done   chan bool
}

func (w worker[I, O]) run(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-w.done:
			return
		case input, ok := <-w.input:
			if !ok {
				return
			}
			fmt.Printf("Worker %d has started\n", w.id)
			output := w.job(input)
			fmt.Printf("Worker %d has ended\n", w.id)
			w.output <- output
		}
	}
}
