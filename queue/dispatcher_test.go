package queue

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestDispatcherJobSimple struct{}

func (j *TestDispatcherJobSimple) Start(finished chan<- interface{}) {
	finished <- true
}

func Test_Dispatcher(t *testing.T) {

	dispatcher := NewDispatcher(runtime.NumCPU(), 10)
	dispatcher.Start()
	stop := false

	for i := 0; i < 100; i++ {
		go dispatcher.Queue(&TestDispatcherJobSimple{})
	}

	stopTimer := time.NewTimer(10 * time.Millisecond)

	for {
		select {
		case <-stopTimer.C:
			stop = true
			dispatcher.Stop()
			break
		}

		if stop {
			break
		}

	}

	assert.EqualValues(t, 100, dispatcher.TotalDispatchedJobs())

}
