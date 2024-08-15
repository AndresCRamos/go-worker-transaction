package worker_pool

import (
	"sync"
)

type WorkerPool[I, O any] struct {
	job        func(I) O
	numWorkers int
	workerList []worker[I, O]
	input      chan I
	output     chan O
	done       chan bool
	wg         *sync.WaitGroup
}

func NewWorkerPool[I, O any](job func(I) O, workers int) *WorkerPool[I, O] {
	input := make(chan I, workers)
	output := make(chan O, workers)

	return &WorkerPool[I, O]{
		job:        job,
		numWorkers: workers,
		input:      input,
		output:     output,
		wg:         new(sync.WaitGroup),
		done:       make(chan bool),
	}
}

func (wk *WorkerPool[I, O]) GetInputChannel() chan<- I {
	return wk.input
}

func (wk *WorkerPool[I, O]) Run() <-chan O {
	// Start workers
	wk.wg.Add(wk.numWorkers)
	for id := 0; id < wk.numWorkers; id++ {
		worker := worker[I, O]{id, wk.input, wk.output, wk.job, wk.done}
		wk.workerList = append(wk.workerList, worker)
		go worker.run(wk.wg)
	}

	// Signal completion when done
	go func() {
		wk.wg.Wait()
		close(wk.output)
		close(wk.done)
	}()

	return wk.output
}
