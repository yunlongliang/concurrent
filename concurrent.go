package concurrent

import "sync"

// Concurrent for simple concurrency control
type Concurrent struct {
	channel chan bool
	wait    *sync.WaitGroup
}

// New initializes a concurrent control object
func New(limit int) *Concurrent {
	return &Concurrent{
		wait:    &sync.WaitGroup{},
		channel: make(chan bool, limit),
	}
}

// Add is used to add a task
func (c *Concurrent) Add(n int) {
	c.wait.Add(n)
	for n > 0 {
		n--
		c.channel <- true
	}
}

// Done is used to accomplish a task
func (c *Concurrent) Done() {
	c.wait.Done()
	<-c.channel
}

// Wait is used to wait for all tasks to be completed
func (c *Concurrent) Wait() {
	c.wait.Wait()
}
