/*
Copyright (c) 2018, Nelson Lin. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Package hashmap implements hashmap.
Actually, golang supplies hashmap very well already
IMHO, it's not necessary to wrap it.

This structure is not thread safe


Reference: http://hg.openjdk.java.net/jdk8/jdk8/jdk/file/687fd7c7986d/src/share/classes/java/util/HashMap.java

*/
package hashmap

import (
	"fmt"
)

/*
Map structure definition
*/
type Map struct {
	m map[interface{}]interface{}
}

func init() {
	ansii := `
	.__                  .__                           
	|  |__ _____    _____|  |__   _____ _____  ______  
	|  |  \\__  \  /  ___/  |  \ /     \\__  \ \____ \ 
	|   Y  \/ __ \_\___ \|   Y  \  Y Y  \/ __ \|  |_> >
	|___|  (____  /____  >___|  /__|_|  (____  /   __/ 
		 \/     \/     \/     \/      \/     \/|__|    
	`
	fmt.Println(ansii)
}

/*
New implementation of HashMap
*/
func New() *Map {
	return &Map{make(map[interface{}]interface{})}
}

/*
Put implementation
*/
func (m *Map) Put(k, v interface{}) interface{} {
	m.m[k] = v
	return v
}

/*
PutIfAbsent states that if the specified key is not already associated with a value (or is mapped to null) associates it with the given value and returns null, else returns the current value.
*/
func (m *Map) PutIfAbsent(k, v interface{}) interface{} {
	if _, ok := m.m[k]; !ok {
		m.m[k] = v
	}
	return m.m[k]
}

/*
Get implementation
*/
func (m *Map) Get(key interface{}) (interface{}, bool) {
	v, ok := m.m[key]
	return v, ok
}

/*
GetOrDefault returns the value to which the specified key is mapped, or defaultValue if this map contains no mapping for the key
*/
func (m *Map) GetOrDefault(key, v interface{}) interface{} {
	if _, ok := m.m[key]; !ok {
		m.m[key] = v
	}
	return m.m[key]
}

/*
Replace implementation
*/
func (m *Map) Replace(k, v, new interface{}) bool {
	if val, ok := m.m[k]; ok {
		if val == v {
			m.m[k] = new
			return true
		}

	}
	return false
}

/*
Remove implementation
*/
func (m *Map) Remove(k interface{}) {
	if _, ok := m.m[k]; ok {
		delete(m.m, k)
	}
}

/*
ContainsKey implementation
*/
func (m *Map) ContainsKey(key interface{}) bool {
	_, ok := m.m[key]
	return ok
}

/*
Size of hashmap, which means number of elements in hashmap
*/
func (m *Map) Size() int {
	return len(m.m)
}

/*
IsEmpty implementation
*/
func (m *Map) IsEmpty() bool {
	return m.Size() == 0
}

/*
PutAll implementations of all of the mappings from the specified map to this map.
*/
func (m *Map) PutAll(other *Map) {
	for k, v := range other.m {
		m.PutIfAbsent(k, v)
	}
}

/*
Keys shows all keys in hashmap
*/
func (m *Map) Keys() []interface{} {
	keys := make([]interface{}, m.Size())
	count := 0
	for k := range m.m {
		keys[count] = k
		count++
	}
	return keys
}

/*
Values shows all values in hashmap
*/
func (m *Map) Values() []interface{} {
	values := make([]interface{}, m.Size())
	count := 0
	for _, v := range m.m {
		values[count] = v
		count++
	}
	return values
}

/*
Clear up the hashmap.
Let GC do it
*/
func (m *Map) Clear() {
	m.m = make(map[interface{}]interface{})
}
