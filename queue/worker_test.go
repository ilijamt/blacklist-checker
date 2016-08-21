package queue

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestWorkerJobSimple struct{}

func (j *TestWorkerJobSimple) Start(finished chan<- interface{}) {
	finished <- true
}

func Test_NewWorker(t *testing.T) {

	dispatchQueue := make(chan Job)
	worker := NewWorker(dispatchQueue)
	assert.NotNil(t, worker)

	stopWorkerTimer := time.NewTimer(11 * time.Millisecond)
	stopCheckIfStopped := time.NewTicker(1 * time.Millisecond)
	addJobsWorkerTicker := time.NewTicker(2 * time.Millisecond)

	go worker.Start()

	for {
		select {
		case <-addJobsWorkerTicker.C:
			dispatchQueue <- &TestWorkerJobSimple{}
		case <-stopWorkerTimer.C:
			stopWorkerTimer.Stop()
			worker.Stop()
		case <-stopCheckIfStopped.C:
		}

		if worker.IsStopped() {
			stopCheckIfStopped.Stop()
			addJobsWorkerTicker.Stop()
			break
		}

	}

	assert.EqualValues(t, 5, worker.ProcessedJobs())
	assert.EqualValues(t, 0, worker.ActiveJobs())

}
