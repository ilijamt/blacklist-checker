package atomic

import "sync"

type Bool struct {
	sync.Mutex
	value bool
}

func (c *Bool) Value() bool {
	c.Lock()
	defer c.Unlock()
	return c.value
}

func (c *Bool) True() {
	c.Lock()
	defer c.Unlock()
	c.value = true
}

func (c *Bool) False() {
	c.Lock()
	defer c.Unlock()
	c.value = false
}
