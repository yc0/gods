/**
* Copyright (c) 2018, Nelson Lin. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
*
* Package arraylist implements the array list.
*
* reference : http://developer.classpath.org/doc/java/util/LinkedList-source.html
*             https://github.com/emirpasic/gods/blob/master/lists/doublylinkedlist/doublylinkedlist.go
 */
package linkedlist

import (
	"errors"
	"fmt"
	"strings"

	"github.com/yc0/gods/utils"
)

type Entry struct {
	data     interface{}
	next     *Entry
	previous *Entry
}

type LinkedList struct {
	first *Entry
	last  *Entry
	size  int
}

func init() {
	anscii := `
	.____    .__        __              .___.____    .__          __   
	|    |   |__| ____ |  | __ ____   __| _/|    |   |__| _______/  |_ 
	|    |   |  |/    \|  |/ // __ \ / __ | |    |   |  |/  ___/\   __\
	|    |___|  |   |  \    <\  ___// /_/ | |    |___|  |\___ \  |  |  
	|_______ \__|___|  /__|_ \\___  >____ | |_______ \__/____  > |__|  
			\/       \/     \/    \/     \/         \/       \/       
	`
	fmt.Println(anscii)
}
func New() *LinkedList {
	return &LinkedList{}
}
func (l *LinkedList) Size() int {
	return l.size
}
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

/*
Returns the first element in the list.
*/
func (l *LinkedList) GetFirst() (interface{}, error) {
	if l.size == 0 {
		return nil, errors.New("NoSuchElementException")
	}
	return l.first.data, nil
}

/*
Returns the last element in the list.
*/
func (l *LinkedList) GetLast() (interface{}, error) {
	if l.size == 0 {
		return nil, errors.New("NoSuchElementException")
	}
	return l.last.data, nil
}

/*
Remove and return the first element in the list.
*/
func (l *LinkedList) RemoveFirst() (interface{}, error) {
	if l.size == 0 {
		return nil, errors.New("NoSuchElementException")
	}
	l.size--
	rst := l.first.data
	if l.first.next == nil {
		l.last = nil
	} else {
		l.first.next.previous = nil
	}
	l.first = l.first.next
	return rst, nil
}

/*
Remove and return the last element in the list.
*/
func (l *LinkedList) RemoveLast() (interface{}, error) {
	if l.size == 0 {
		return nil, errors.New("NoSuchElementException")
	}
	l.size--
	rst := l.last.data
	if l.last.previous == nil {
		l.first = nil
	} else {
		l.last.previous.next = nil
	}
	l.last = l.last.previous
	return rst, nil
}

/*
Insert an element at the first of the list.
*/

func (l *LinkedList) AddFirst(data interface{}) {
	e := &Entry{data, nil, nil}
	if l.size == 0 {
		l.first = e
		l.last = e

	} else {
		e.next = l.first
		l.first.previous = e
		l.first = e
	}
	l.size++
}

func (l *LinkedList) Prepend(data interface{}) {
	l.AddFirst(data)
}

/*
Insert an element at the last of the list.
*/

func (l *LinkedList) AddLast(data interface{}) {
	e := &Entry{data, nil, nil}
	if l.size == 0 {
		l.first = e
		l.last = e
	} else {
		l.last.next = e
		e.previous = l.last
		l.last = e
	}
	l.size++
}
func (l *LinkedList) Append(data interface{}) {
	l.AddLast(data)
}
func (l *LinkedList) Contains(value interface{}) bool {
	if l.size == 0 {
		return false
	}
	for cur := l.first; cur != nil; cur = cur.next {
		if cur.data == value {
			return true
		}
	}
	return false
}

func (l *LinkedList) Get(index int) (interface{}, error) {
	e, err := l.getEntry(index)
	if err != nil {
		return nil, err
	}
	return e.data, nil
}

func (l *LinkedList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= l.size {
		return nil, errors.New("out of bound")
	}
	e, err := l.getEntry(index)
	if err != nil {
		return nil, err
	}
	l.size--
	if l.size == 0 {
		l.first = nil
		l.last = nil
	} else if e == l.first {
		l.first = e.next
		e.next.previous = nil
	} else if e == l.last {
		e.previous.next = nil
		l.last = e.previous
	} else {
		e.previous.next = e.next
		e.next.previous = e.previous
	}
	return e.data, nil
}

func (l *LinkedList) Add(values ...interface{}) {
	for _, v := range values {
		l.AddLast(v)
	}
}
func (l *LinkedList) AddAt(index int, value interface{}) (int, error) {
	e, err := l.getEntry(index)
	if err != nil {
		return -1, err
	}
	l.size++
	entry := &Entry{value, nil, nil}
	entry.next = e
	if e == l.first {
		e.previous = entry
		l.first = entry
	} else {
		entry.previous = e.previous
		e.previous.next = entry
		e.previous = entry
		// if e == l.last {
		// 	l.last = entry
		// }
	}
	return index, nil
}
func (l *LinkedList) IndexOf(value interface{}) int {
	if l.size == 0 {
		return -1
	}
	for i, cur := 0, l.first; cur != nil; i, cur = i+1, cur.next {
		if cur.data == value {
			return i
		}
	}
	return -1
}
func (l *LinkedList) Clear() {
	l.size = 0
	l.first = nil
	l.last = nil
}
func (l *LinkedList) Sort() {
	values := l.values()
	utils.Sort(values)
	l.Clear()
	l.Add(values)
}

func (l *LinkedList) Clone() *LinkedList {
	new_linkedlist := New()

	values := l.values()
	new_linkedlist.Add(values...)
	return new_linkedlist
}

func (l *LinkedList) String() string {
	elements := []string{}
	for cur := l.first; cur != nil; cur = cur.next {
		elements = append(elements, fmt.Sprintf("%v", cur.data))
	}
	return fmt.Sprintf("&LinkedList{ %s }", strings.Join(elements, ","))
}

// internal functions

func (l *LinkedList) values() []interface{} {
	values := make([]interface{}, l.size)
	for i, cur := 0, l.first; cur != nil; i, cur = i+1, cur.next {
		values[i] = cur.data
	}
	return values
}

func (l *LinkedList) getEntry(index int) (*Entry, error) {
	if index < 0 || index >= l.size {
		return nil, errors.New("out of bound")
	}
	// starts with first
	var e *Entry
	if index < l.size/2 {
		e = l.first
		for index > 0 {
			e = e.next
			index--
		}
	} else {
		e = l.last
		for i := index + 1; i < l.size; i++ {
			e = e.previous
		}
	}
	return e, nil
}
