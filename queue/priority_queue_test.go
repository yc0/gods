package queue

import (
	"testing"
)

func TestQueueAppend(t *testing.T) {
	pq := New()
	t.Logf("%+v", pq)
}
