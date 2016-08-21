package queue

import "github.com/ilijamt/blacklist-checker/atomic"

type Dispatcher struct {
	workers map[string]*Worker

	totalWorkers int

	queue chan Job
	stop  *atomic.Bool

	jobsDispatched *atomic.IntNumber
}

func NewDispatcher(totalWorkers int, queueSize int) *Dispatcher {
	return &Dispatcher{
		totalWorkers:   totalWorkers,
		workers:        make(map[string]*Worker),
		queue:          make(chan Job, queueSize),
		stop:           &atomic.Bool{},
		jobsDispatched: &atomic.IntNumber{},
	}
}

func (d Dispatcher) Start() {
	go func() {
		for i := 1; i <= d.totalWorkers; i++ {
			worker := NewWorker(d.queue)
			worker.Start()
			d.workers[worker.ID] = worker
		}
	}()
}

func (d Dispatcher) TotalDispatchedJobs() int64 {
	return d.jobsDispatched.Value()
}

func (d Dispatcher) Stop() {
	d.stop.True()
	for id := range d.workers {
		d.workers[id].Stop()
	}
}

func (d Dispatcher) Queue(j Job) {
	if !d.stop.Value() {
		d.queue <- j
		d.jobsDispatched.Incr()
	}
}
