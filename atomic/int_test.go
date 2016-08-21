package atomic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Int(t *testing.T) {

	var finalValue int64 = 10
	obj := IntNumber{}

	obj.Incr()
	assert.EqualValues(t, 1, obj.Value())
	obj.Decr()
	assert.EqualValues(t, 0, obj.Value())

	var i int64
	for i = 0; i <= finalValue; i++ {
		obj.Set(i)
	}

	assert.EqualValues(t, finalValue, obj.Value())

}
