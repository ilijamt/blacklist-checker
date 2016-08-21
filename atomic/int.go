package atomic

import "sync"

type IntNumber struct {
	sync.Mutex
	value int64
}

func (c *IntNumber) Set(value int64) {
	c.Lock()
	defer c.Unlock()
	c.value = value
}

func (c *IntNumber) Value() int64 {
	c.Lock()
	defer c.Unlock()
	return c.value
}

func (c *IntNumber) Add(value int64) {
	c.Lock()
	defer c.Unlock()
	c.value += value
}

func (c *IntNumber) Remove(value int64) {
	c.Lock()
	defer c.Unlock()
	c.value -= value
}

func (c *IntNumber) Incr() {
	c.Add(1)
}

func (c *IntNumber) Decr() {
	c.Remove(1)
}
