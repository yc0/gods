/*
Copyright (c) 2018, Nelson Lin. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

The List interface claim protocols here.

Structure is not thread safe,and remember that array in golang is call by value.
It costs a lot while you try to pass as an argument. Therefore, we ought to handle it
carefully.
*/
package lists

type List interface {
	New() *List
	Clone() *List
	Get(index int) (interface{}, error)
	Remove(index int) (interface{}, error)
	Add(values ...interface{})
	Contains(value interface{}) bool
	IndexOf(value interface{}) int
	Sort()
	Clear()
	Size() int
	String() string
	AddAt(index int, value interface{}) (int, error)
}
