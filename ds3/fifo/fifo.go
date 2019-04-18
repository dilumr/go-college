package fifo

import (
	"container/list"
)

type FifoQueue struct {
	underlying *list.List
}

func New() *FifoQueue {
	return &FifoQueue{list.New()}
}

func (fifo *FifoQueue) Peek() (interface{}, bool) {
	el := fifo.underlying.Front()
	if el == nil {
		return nil, false
	}
	return el.Value, true
}

func (fifo *FifoQueue) Dequeue() (interface{}, bool) {
	el := fifo.underlying.Front()
	if el == nil {
		return nil, false
	}
	fifo.underlying.Remove(el)
	return el.Value, true
}

func (fifo *FifoQueue) Enqueue(value interface{}) {
	fifo.underlying.PushBack(value)
}
