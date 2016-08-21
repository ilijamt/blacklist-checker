package atomic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Bool(t *testing.T) {

	obj := Bool{}

	obj.True()
	assert.True(t, obj.Value())
	obj.False()
	assert.False(t, obj.Value())

}
