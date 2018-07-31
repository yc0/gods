/*
Package maps provides an abstract Map interface.

Copyright (c) 2018, Nelson Lin. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
*/
package maps

/*
Map associate with array, map, symbol table or dictionary
*/
type Map interface {
	Put(key interface{}, value interface{})
	Get(key interface{}) (value interface{}, found bool)
	Remove(key interface{})
	Keys() []interface{}
}
