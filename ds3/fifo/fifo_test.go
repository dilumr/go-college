package fifo

import (
	"fmt"
	"testing"
)

func TestFifoQueue(t *testing.T) {
	q := New()
	for i := 0; i < 100; i++ {
		q.Enqueue(i)
	}
	for v, _ := q.Dequeue(); v != nil; v, _ = q.Dequeue() {
		fmt.Println(v)
	}
}
