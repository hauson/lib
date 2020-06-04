package exchan

import (
	"github.com/lib/container/queue"
	"github.com/lib/concurrence/trylock"
)

// ExChan wrapper chan with infinite capacity
type ExChan struct {
	lock  trylock.TryLock
	queue *queue.Queue
	in    chan interface{}
	out   chan interface{}
	exit  chan int
}

func New() *ExChan {
	c := &ExChan{
		lock:  trylock.New(),
		queue: queue.New(),
		in:    make(chan interface{}, 1024),
		out:   make(chan interface{}, 1024),
		exit:  make(chan int),
	}

	go c.run()
	return c
}

func (c *ExChan) run() {
	var outData interface{}
	for {
		if c.queue.Empty() {
			select {
			case inData := <-c.in:
				c.queue.EnQueue(inData)
			case <-c.exit:
				return
			}
		} else {
			if c.lock.TryLock() {
				outData = c.queue.DeQueue()
			}

			select {
			case inData := <-c.in:
				c.queue.EnQueue(inData)
			case c.out <- outData:
				c.lock.UnLock()
			case <-c.exit:
				return
			}
		}
	}
}

func (c *ExChan) OutCh() <-chan interface{} {
	return c.out
}

func (c *ExChan) InCh() chan<- interface{} {
	return c.in
}

func (c *ExChan) Close() {
	close(c.exit)
}
