package queue

import (
	"testing"
)

func TestQueueAppendPoll(t *testing.T) {
	pq := New()

	pq.Offer(5, 8, 1, 2, 3, 4, 9, 10, -1)

	if v := pq.Peek(); v != -1 {
		t.Errorf("peek expect:%v, but %v", -1, v)
	}
	prev := pq.Poll()
	for i := 1; i < pq.Size(); i++ {
		c := pq.Poll()
		if c.(int) < prev.(int) {
			t.Errorf("no sifting up of priority queue. prev(%v), current(%v)", prev, c)
		}
		prev = c
	}

}
