/*
Copyright (c) 2018, Nelson Lin. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Package queue currently implements the priorityqueue and queue interface.

The priorityqueue essentially bases on bubbling up for sorting in-place, where the way is meant to heap-sort.
Therefore, bubble up take O(logn)

reference: http://hg.openjdk.java.net/jdk8/jdk8/jdk/file/687fd7c7986d/src/share/classes/java/util/PriorityQueue.java

*/
package queue

import (
	"errors"
	"fmt"

	"github.com/yc0/gods/utils"
)

func init() {
	ansii := `
	__________        .__             .__  __          ________                               
	\______   \_______|__| ___________|__|/  |_ ___.__.\_____  \  __ __   ____  __ __   ____  
	 |     ___/\_  __ \  |/  _ \_  __ \  \   __<   |  | /  / \  \|  |  \_/ __ \|  |  \_/ __ \ 
	 |    |     |  | \/  (  <_> )  | \/  ||  |  \___  |/   \_/.  \  |  /\  ___/|  |  /\  ___/ 
	 |____|     |__|  |__|\____/|__|  |__||__|  / ____|\_____\ \_/____/  \___  >____/  \___  >
	                                            \/            \__>           \/            \/ 
	`
	fmt.Println(ansii)
}

/*
PriorityQueue struct
*/
type PriorityQueue struct {
	queue []interface{}
	size  int
}

/*
New PriorityQueue
*/
func New() *PriorityQueue {
	return &PriorityQueue{make([]interface{}, 11), 0}
}

/*
NewSlice denotes new a priorityqueue with a slice
*/
func NewSlice(a ...interface{}) *PriorityQueue {
	size := len(a)
	if size < 11 {
		size = 11
	}
	slice := make([]interface{}, size)
	copy(slice, a)
	pq := &PriorityQueue{slice, len(a)}
	fmt.Printf("+%v", pq)
	pq.heapify()
	return pq
}

/*
Clear removes all elements from the list.
Here, we let GC do it.
The queue will be empty after this call returns
*/
func (pq *PriorityQueue) Clear() {
	pq.queue = make([]interface{}, 11)
	pq.size = 0
}

/*
Offer the specified element into this priority queue
*/
func (pq *PriorityQueue) Offer(o ...interface{}) (bool, error) {
	if o == nil {
		return false, errors.New("NullPointerException")
	}
	pq.grow(len(o))
	for _, e := range o {
		pq.queue[pq.size] = e
		pq.siftUp(pq.size, e)
		pq.size++
	}
	return true, nil
}

/*
Add the specified element into this priority queue alternatively
*/
func (pq *PriorityQueue) Add(o ...interface{}) (bool, error) {
	return pq.Offer(o...)
}

/*
Poll queue, and represent the smallest in the queue.
Then, sift down the bigger one to recliam smallest one on the top again
*/
func (pq *PriorityQueue) Poll() interface{} {
	if pq.size == 0 {
		return nil
	}
	pq.size--
	rst := pq.queue[0]
	e := pq.queue[pq.size]
	pq.queue[pq.size] = nil
	if pq.size != 0 {
		pq.siftDown(0, e)
	}
	return rst
}

/*
Peek first element of queue. In other words, no poll
element of queue, but take a look
*/
func (pq *PriorityQueue) Peek() interface{} {
	if pq.size == 0 {
		return nil
	}
	return pq.queue[0]
}

/*
Remove a single instance of the specified element from this queue,
if it is present.
*/
func (pq *PriorityQueue) Remove(o interface{}) bool {
	idx := pq.indexOf(o)
	if idx == -1 {
		return false
	}
	pq.removeAt(idx)
	return true
}

/*
Contains the specified element or not
*/
func (pq *PriorityQueue) Contains(o interface{}) bool {
	return pq.indexOf(o) != -1
}

func (pq *PriorityQueue) indexOf(o interface{}) int {
	if o != nil {
		for i, v := range pq.queue {
			if v == o {
				return i
			}
		}
	}
	return -1
}

/**
 * Removes the ith element from queue.
 *
 * Normally this method leaves the elements at up to i-1,
 * inclusive, untouched.  Under these circumstances, it returns
 * null.  Occasionally, in order to maintain the heap invariant,
 * it must swap a later element of the list with one earlier than
 * i.  Under these circumstances, this method returns the element
 * that was previously at the end of the list and is now at some
 * position before i. This fact is used by iterator.remove so as to
 * avoid missing traversing elements.
 */
func (pq *PriorityQueue) removeAt(idx int) {
	pq.size--
	rst := pq.queue[pq.size]
	pq.queue[pq.size] = nil
	if idx == pq.size {
		return
	}
	pq.siftDown(idx, rst)
	if pq.queue[idx] == rst {
		// bubble up util idx
		// so we keep checking if it can be bubbled up
		pq.siftUp(idx, rst)

	}
}
func (pq *PriorityQueue) siftUp(k int, x interface{}) {
	// while (k > 0) {
	// 	int parent = (k - 1) >>> 1;
	// 	Object e = queue[parent];
	// 	if (comparator.compare(x, (E) e) >= 0)
	// 		break;
	// 	queue[k] = e;
	// 	k = parent;
	// }
	// queue[k] = x;
	for k > 0 {
		// here, we need do unsigned shift
		parent := uint(k-1) >> 1
		e := pq.queue[parent]
		comparator := utils.GetComparator(e)
		if comparator(x, e) >= 0 {
			break
		}
		pq.queue[k] = e
		k = int(parent)
	}
	pq.queue[k] = x
}

func (pq *PriorityQueue) grow(n int) {
	oldCapacity := cap(pq.queue)
	if pq.size+n >= oldCapacity {
		// Double size if small; else grow by 50%
		var newCapacity int
		if oldCapacity < 64 {
			newCapacity = oldCapacity + oldCapacity + 2
		} else {
			newCapacity = oldCapacity + (oldCapacity >> 1)
		}
		if newCapacity-n < 0 {
			newCapacity = n
		}
		pq.resize(newCapacity)
	}
}

func (pq *PriorityQueue) resize(cap int) {
	newObjects := make([]interface{}, cap)
	copy(newObjects, pq.queue)
	pq.queue = newObjects
}

func (pq *PriorityQueue) heapify() {
	for i := (pq.size>>1 - 1); i >= 0; i-- {
		pq.siftDown(i, pq.queue[i])
	}
}

func (pq *PriorityQueue) siftDown(k int, x interface{}) {
	half := uint(pq.size) >> 1 // loop while a none-leaf
	for k < int(half) {
		child := k<<1 + 1
		right := child + 1
		c := pq.queue[child]
		comparator := utils.GetComparator(c)

		// find left and right trees whose values is the smallest
		if right < pq.size && comparator(c, pq.queue[right]) > 0 {
			// right is smaller
			c, child = pq.queue[right], right
		}
		if comparator(x, c) <= 0 {
			// current parent is the smallest
			break
		}
		// otherwise, go through down
		pq.queue[k] = c
		k = child
	}
	pq.queue[k] = x
}

/*
Size is shown
*/
func (pq *PriorityQueue) Size() int {
	return pq.size
}
