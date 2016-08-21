package queue

import (
	"github.com/satori/go.uuid"

	"github.com/ilijamt/blacklist-checker/atomic"
)

type Worker struct {
	ID string

	queue       <-chan Job
	stop        chan bool
	jobFinished chan interface{}

	active    *atomic.IntNumber
	processed *atomic.IntNumber
	stopped   *atomic.Bool
}

func NewWorker(workQueue chan Job) *Worker {
	return &Worker{
		ID: uuid.NewV4().String(),

		queue: workQueue,
		stop:  make(chan bool),

		jobFinished: make(chan interface{}),

		stopped: &atomic.Bool{},

		active:    &atomic.IntNumber{},
		processed: &atomic.IntNumber{},
	}
}

func (w *Worker) ActiveJobs() int64 {
	return w.active.Value()
}

func (w *Worker) ProcessedJobs() int64 {
	return w.processed.Value()
}

func (w *Worker) IsStopped() bool {
	return w.stopped.Value()
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.queue:
				go func() {
					w.active.Incr()
					job.Start(w.jobFinished)
				}()
			case <-w.jobFinished:
				go func() {
					w.active.Decr()
					w.processed.Incr()
				}()
			case <-w.stop:
				w.stopped.True()
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.stop <- true
	}()
}
