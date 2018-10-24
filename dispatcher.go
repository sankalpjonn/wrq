package wrq

import (
  "sync"
)

const (
  DEFAULT_NAME       = "wrq"
  DEFAULT_QUEUE_SIZE = 100
  DEFAULT_WORKERS    = 100
)

type Dispatcher struct {
	name       string
	workerPool chan chan Job
	maxWorkers int
	jobQueue   chan Job
	wg         *sync.WaitGroup
	doneCh     chan bool
}

func New() *Dispatcher {
  return NewWithSettings(DEFAULT_NAME, DEFAULT_QUEUE_SIZE, DEFAULT_WORKERS)
}

func NewWithSettings(name string, queueSize int, maxWorkers int) *Dispatcher {
  workerPool := make(chan chan Job, maxWorkers)
	jobQueue := make(chan Job, queueSize)

	return &Dispatcher{
		name:       name,
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		workerPool: workerPool,
		wg:         &sync.WaitGroup{},
		doneCh:     make(chan bool),
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		id := i + 1
		d.wg.Add(1)
		worker := NewWorker(id, d.workerPool, d.wg)
		worker.start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for job := range d.jobQueue {
		workerJobQueue := <-d.workerPool
		workerJobQueue <- job
	}

	for {
		select {
		case worker, ok := <-d.workerPool:
			if ok {
				close(worker)
			} else {
				d.doneCh <- true
				return
			}

		}
	}

}

func (d *Dispatcher) AddJob(job Job) {
	d.jobQueue <- job
}

func (d *Dispatcher) Stop() {
	// No more Adding jobs to the jobqueue function
	close(d.jobQueue)
	d.wg.Wait()
	close(d.workerPool)
	<-d.doneCh
}
