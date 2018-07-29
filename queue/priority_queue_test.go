package queue

import (
	"testing"
)

func TestQueueAppend(t *testing.T) {
	pq := New()
	pq.Offer(5, 8, 1, 2, 3, 4, 9, 10, -1)
	t.Logf("%+v", pq)
}
