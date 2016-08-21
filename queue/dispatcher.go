package queue

type Dispatcher struct {
	workers map[string]*Worker

	totalWorkers int

	queue chan Job
}

func NewDispatcher(totalWorkers int, queueSize int) *Dispatcher {
	return &Dispatcher{
		totalWorkers: totalWorkers,
		workers:      make(map[string]*Worker),
		queue:        make(chan Job, queueSize),
	}
}

func (d Dispatcher) Start() {
	go func() {
		for i := 1; i <= d.totalWorkers; i++ {
			worker := NewWorker(d.queue)
			d.workers[worker.ID] = worker
		}
	}()
}

func (d Dispatcher) Stop() {
	for id := range d.workers {
		go d.workers[id].Stop()
	}
}

func (d Dispatcher) Queue(j Job) {
	d.queue <- j
}
