/*
Copyright (c) 2018, Nelson Lin. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
*/
package queue

type Queue interface {
	Clear()
	Add(o ...interface{}) (bool, error)
	Offer(o ...interface{}) (bool, error)
	Peek() interface{}
	Poll() interface{}
	Remove(o interface{}) bool
	Size() interface{}
	Contains(o interface{}) bool
}
