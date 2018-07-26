/*
Copyright (c) 2018, Nelson Lin. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Package arraylist implements the array list.

This structure is thread safe


Reference: http://developer.classpath.org/doc/java/util/Vector-source.html

*/
package vector

import (
	"errors"
	"fmt"

	"github.com/yc0/gods/utils"
)

type Vector struct {
	objects []interface{}
	size    int
}

const (
	MAX_ARRAY_SIZE = 1<<31 - 4
)

func init() {
	anscii := `
   \   \ /   /____   _____/  |_  ___________ 
	\   Y   // __ \_/ ___\   __\/  _ \_  __ \
	 \     /\  ___/\  \___|  | (  <_> )  | \/
	  \___/  \___  >\___  >__|  \____/|__|   
				 \/     \/                       
	`
	fmt.Println(anscii)
}

func New() *Vector {
	return &Vector{make([]interface{}, 10), 0}
}

func (l *Vector) Add(values ...interface{}) {
	l.grow(len(values))
	for _, val := range values {
		l.objects[l.size] = val
		l.size++
	}
}

func (l *Vector) AddAt(index int, obj interface{}) (int, error) {
	if index < 0 || index > l.size {
		return -1, errors.New("out of bound")
	}
	l.grow(1)
	// java here adopt System.arraycopy(src, src_idx, dest, dest_idx, len)
	copy(l.objects[index+1:], l.objects[index:])
	l.objects[index] = obj
	l.size++
	return index, nil
}

/**
 * Returns the element at the specified position in this list.
 *
 * @param  index index of the element to return
 * @return (object, error) the element at the specified position in this list
 */
func (l *Vector) Get(idx int) (interface{}, error) {
	if idx < 0 || idx >= l.size {
		return nil, errors.New("IndexOutOfBoundsException")
	}
	return l.objects[idx], nil
}

func (l *Vector) Set(idx int, obj interface{}) (int, error) {
	if idx < 0 || idx >= l.size {
		return -1, errors.New("IndexOutOfBoundsException")
	}
	l.objects[idx] = obj
	return idx, nil
}

/**
 * Returns true if this list contains the specified element
 */
func (l *Vector) Contains(obj interface{}) bool {
	return l.IndexOf(obj) >= 0
}

/**
 * Returns the index of the "first" occurrence of the specified element
 * in this list, or -1 if this list does not contain the element.
 */
func (l *Vector) IndexOf(obj interface{}) int {
	if l.size == 0 {
		return -1
	}

	for i, o := range l.objects {
		if o == obj {
			return i
		}
	}
	return -1
}

func (l *Vector) Clone() *Vector {
	newone := New()
	newone.objects = make([]interface{}, len(l.objects))
	copy(newone.objects, l.objects)
	newone.size = l.size
	return newone
}

/**
 * We left this implementation. We can manipulate native sort way
 * , and need to implement three methods.

 * type Interface interface {
 * 		// Len is the number of elements in the collection.
 * 		Len() int
 * 		// Less reports whether the element with
 * 		// index i should sort before the element with index j.
 * 		Less(i, j int) bool
 * 		// Swap swaps the elements with indexes i and j.
 * 		Swap(i, j int)
 * }


 * Here is implementations in Java
 * public void sort(Comparator<? super E> c) {
 *    final int expectedModCount = modCount;
 *    Arrays.sort((E[]) elementData, 0, size, c);
 *    if (modCpl[.gyfcount != expectedModCount) {
 *        throw new ConcurrentModificationException();
 *    }
 *     modCount++;
 * }
 */
func (l *Vector) Sort() {
	utils.Sort(l.objects[:l.size]) // this way would provide high performance by constrained slice
}

/**
 * Clear removes all elements from the list.
 * Here, we let GC do it
 */
func (l *Vector) Clear() {
	l.objects = make([]interface{}, 10)
	l.size = 0
}

func (l *Vector) Remove(idx int) (interface{}, error) {
	if idx < 0 || idx >= l.size {
		return nil, errors.New("IndexOutOfBoundsException")
	}
	var rst interface{}
	rst = l.objects[idx]
	l.objects[idx] = nil
	copy(l.objects[idx:], l.objects[idx+1:l.size])
	l.objects[l.size-1] = nil
	l.size--

	return rst, nil
}

/**
 * Size returns number of elements within the list.
 */
func (l *Vector) Size() int {
	return l.size
}

func (l *Vector) IsEmpty() bool {
	return l.size == 0
}

func (l *Vector) resize(cap int) {
	newObjects := make([]interface{}, cap, cap)
	copy(newObjects, l.objects)
	l.objects = newObjects
}

func (list *Vector) Insert(index int, values ...interface{}) {

	l := len(values)
	list.grow(l)
	list.size += l
	// Shift old to right
	for i := list.size - 1; i >= index+l; i-- {
		list.objects[i] = list.objects[i-l]
	}
	// Insert new
	for i, value := range values {
		list.objects[index+i] = value
	}
}

/**
 * Increases the capacity to ensure that it can hold at least the
 * number of elements specified by the minimum capacity argument.
 *
 * @param n the desired minimum capacity
 */
func (l *Vector) grow(n int) {
	oldCapacity := cap(l.objects)
	if l.size+n >= oldCapacity {
		newCapacity := oldCapacity + (oldCapacity >> 1)
		if newCapacity-n < 0 {
			newCapacity = n
		}
		l.resize(newCapacity)
	}
}
