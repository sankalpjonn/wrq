package wrq

import (
	"sync"
)

type Worker struct {
	id         int
	jobQueue   chan Job
	workerPool chan chan Job
	wg         *sync.WaitGroup
}

func NewWorker(id int, workerPool chan chan Job, wg *sync.WaitGroup) *Worker {
	return &Worker{
		id:         id,
		jobQueue:   make(chan Job),
		workerPool: workerPool,
		wg:         wg,
	}
}

func (w *Worker) start() {
	// defer w.wg.Done()
	go func() {
		defer func() {
			w.wg.Done()
		}()

		w.workerPool <- w.jobQueue

		for job := range w.jobQueue {
			w.workerPool <- w.jobQueue
			job.Execute()
		}
	}()
}

func (w *Worker) close() {
	close(w.jobQueue)
}
