package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewWorker(t *testing.T) {

	dispatchQueue := make(chan Job, 10)
	worker := NewWorker(dispatchQueue)
	assert.NotNil(t, worker)

}
