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
	pq.Add(-2, 6)
	if rst := pq.Peek(); rst != -2 {
		t.Errorf("expect get minimum item -2 but %v", rst)
	}
}

func TestQueueContains(t *testing.T) {
	pq := New()

	pq.Offer(5, 8, 1, 2, 3, 4, 9, 10, -1)

	if rst := pq.Contains(-1); rst == false {
		t.Errorf("contain -1 should be %v, but %v", true, false)
	}
}

func TestQueueClear(t *testing.T) {
	pq := New()

	pq.Offer(5, 8, 1, 2, 3, 4, 9, 10, -1)
	pq.Clear()
	if rst := pq.Size(); rst > 0 {
		t.Errorf("Except 0 but get %v", rst)
	}
}

func TestQueueRemove(t *testing.T) {
	pq := New()
	pq.Offer(5, 8, 1, 2, 3, 4, 9, 10, -1)
	if rst := pq.Remove(-1); rst != true {
		t.Errorf("Removing out -1 should be true, but false")
	}
	if rst := pq.Remove(-1); rst != false {
		t.Errorf("Removing out -1 while queue doesn't exist -1 should be false, but true")
	}
	// priority queue ought to keep heap tree while removing any item.
	size := pq.Size()
	k, half := 0, size>>1
	for k < half {
		left := k<<1 + 1
		right := left + 1
		if right < size && pq.queue[right].(int) < pq.queue[k].(int) {
			t.Errorf("expect right child(%v) is bigger than its parent(%v)", pq.queue[right], pq.queue[k])
		}
		if pq.queue[left].(int) < pq.queue[k].(int) {
			t.Errorf("expect left child(%v) is bigger than its parent(%v)", pq.queue[left], pq.queue[k])
		}
		k = left
	}
}

func TestQueueNewSlice(t *testing.T) {
	a := []interface{}{9, 8, 7, 6, 5, 4, 3, 2, -1, 0}
	pq := NewSlice(a...)
	size := pq.Size()
	k, half := 0, size>>1
	for k < half {
		left := k<<1 + 1
		right := left + 1
		if right < size && pq.queue[right].(int) < pq.queue[k].(int) {
			t.Errorf("expect right child(%v) is bigger than its parent(%v)", pq.queue[right], pq.queue[k])
		}
		if pq.queue[left].(int) < pq.queue[k].(int) {
			t.Errorf("expect left child(%v) is bigger than its parent(%v)", pq.queue[left], pq.queue[k])
		}
		k = left
	}
	k = 0
	for k < half {
		left := k<<1 + 1
		right := left + 1
		if right < size && pq.queue[right].(int) < pq.queue[k].(int) {
			t.Errorf("expect right child(%v) is bigger than its parent(%v)", pq.queue[right], pq.queue[k])
		}
		if pq.queue[left].(int) < pq.queue[k].(int) {
			t.Errorf("expect left child(%v) is bigger than its parent(%v)", pq.queue[left], pq.queue[k])
		}
		k = right
	}
}
