/*
Copyright (c) 2018, Nelson Lin. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Package arraylist implements the array list.

Structure is not thread safe,and remember that array in golang is call by value.
It costs a lot while you try to pass as an argument. Therefore, we ought to handle it
carefully.

If you want to find thread safe structure. Try goto vector.


Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
		   https://github.com/emirpasic/gods/blob/master/lists/arraylist/arraylist.go

*/
package arraylist

import (
	"errors"
	"fmt"
)

type List struct {
	objects []interface{}
	size    int
}

const (
	MAX_ARRAY_SIZE = 1<<31 - 4
)

func init() {
	anscii := `
	_____                             .____    .__          __   
	/  _  \___________________  ___.__.|    |   |__| _______/  |_ 
   /  /_\  \_  __ \_  __ \__  \<   |  ||    |   |  |/  ___/\   __\
  /    |    \  | \/|  | \// __ \\___  ||    |___|  |\___ \  |  |  
  \____|__  /__|   |__|  (____  / ____||_______ \__/____  > |__|  
		  \/                  \/\/             \/       \/        
	`
	fmt.Println(anscii)
}

func New() *List {
	return &List{}
}

func (l *List) Add(values ...interface{}) {
	l.grow(len(values))
	for _, val := range values {
		l.objects[l.size] = val
		l.size++
	}
}

func (l *List) AddAt(index int, obj interface{}) (int, error) {
	if index < 0 || index > l.size {
		return -1, errors.New("out of bound")
	}
	l.grow(1)
	// java here adopt System.arraycopy(src, src_idx, dest, dest_idx, len)
	copy(l.objects[index+1:], l.objects[index:])
	l.objects[index] = obj
	return index, nil
}
func (l *List) resize(cap int) {
	newObjects := make([]interface{}, cap, cap)
	copy(newObjects, l.objects)
	l.objects = newObjects
}
func (l *List) grow(n int) {
	oldCapacity := cap(l.objects)
	if l.size+n >= oldCapacity {
		newCapacity := oldCapacity + (oldCapacity >> 1)
		if newCapacity-n < 0 {
			newCapacity = n
		}
		l.resize(newCapacity)
	}
}
