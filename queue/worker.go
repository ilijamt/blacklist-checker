package queue

import "github.com/satori/go.uuid"

type Worker struct {
	ID    string
	queue <-chan Job
	stop  chan bool
}

func NewWorker(workQueue chan Job) *Worker {
	return &Worker{
		ID:    uuid.NewV4().String(),
		queue: workQueue,
		stop:  make(chan bool, 1),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.queue:
				job.Start()
			case <-w.stop:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.stop <- true
	}()
}
