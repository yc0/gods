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

type PriorityQueue struct {
	queue []interface{}
	size  int
}

func New() *PriorityQueue {
	return &PriorityQueue{make([]interface{}, 11), 0}
}

func (pq *PriorityQueue) Clear() {

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

func (pq *PriorityQueue) Poll() interface{} {
	return nil
}
func (pq *PriorityQueue) Peek() interface{} {
	return nil
}
func (pq *PriorityQueue) Remove() bool {
	return false
}

func (pq *PriorityQueue) findSlot(start int) int {
	return -1
}

func (pq *PriorityQueue) bubbleUp(index int) {

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

// Offer(o interface{}) bool
// Peek() interface{}
// Poll() interface{}
// Remove() bool
// Size() interface{}
