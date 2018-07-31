package hashmap

import (
	"testing"
)

func TestNew(t *testing.T) {
	m := New()

	nxt := New()
	nxt.Put("k", 1)
	nxt.PutIfAbsent("k", 2)
	if val, _ := nxt.Get("k"); val != 1 {
		t.Errorf("PutIfAbsent expects 1 but %v", val)
	}
	if val := nxt.GetOrDefault("i", 3); val != 3 {
		t.Errorf("GetOrDefault exepcts 3 but %v", val)
	}
	m.PutAll(nxt)
	if val := m.Size(); val != 2 {
		t.Errorf("Putall exepcts 2 but %v", val)
	}
}

func TestKeysValues(t *testing.T) {
	m := New()
	m.PutIfAbsent("a", 'a')
	m.PutIfAbsent("p", 'p')
	m.PutIfAbsent("p", 0) // replicate key validate
	m.PutIfAbsent("l", 'l')
	m.PutIfAbsent("e", 'e')

	keys := m.Keys()
	targets := []string{"a", "p", "l", "e"}
	if len(keys) != len(targets) {
		t.Errorf("keys and targets are not the same")
	}
	for _, k := range keys {
		match := false
		for _, word := range targets {
			if word == k.(string) {
				match = true
				break
			}
		}
		if !match {
			t.Errorf("keys and targets are not the same")
			break
		}
	}

	values := m.Values()
	nTargets := []int{'a', 'p', 'l', 'e'}
	if len(values) != len(nTargets) {
		t.Errorf("values and targets are not the same")
	}
	for _, v := range values {
		match := false
		for _, val := range nTargets {
			if int32(val) == v.(int32) {
				match = true
				break
			}
		}
		if !match {
			t.Errorf("values and targets are not the same")
			break
		}
	}
}
func TestClear(t *testing.T) {
	m := New()
	m.PutIfAbsent("a", 'a')
	m.PutIfAbsent("p", 'p')
	m.PutIfAbsent("p", 0) // replicate key validate
	m.PutIfAbsent("l", 'l')
	m.PutIfAbsent("e", 'e')
	m.Clear()
	if val := m.Size(); val != 0 {
		t.Errorf("Clear expect size:0, but %v", val)
	}
	if val := m.IsEmpty(); val != true {
		t.Errorf("Isempty expects true, but false")
	}
}

func TestContainsKey(t *testing.T) {
	m := New()
	m.PutIfAbsent("a", 'a')
	m.PutIfAbsent("p", 'p')
	m.PutIfAbsent("p", 0) // replicate key validate
	m.PutIfAbsent("l", 'l')
	m.PutIfAbsent("e", 'e')
	if val := m.ContainsKey("o"); val != false {
		t.Errorf("Contain expect false, but true")
	}

	if val := m.ContainsKey("a"); val != true {
		t.Errorf("Contain expect true, but false")
	}
}
func TestRemoveReplace(t *testing.T) {
	m := New()
	m.PutIfAbsent("a", 'a')
	m.PutIfAbsent("p", 'p')
	m.PutIfAbsent("p", 0) // replicate key validate
	m.PutIfAbsent("l", 'l')
	m.PutIfAbsent("e", 'e')

	m.Remove("p")
	if _, ok := m.Get("p"); ok {
		t.Errorf("Remove expect false, but true")
	}

	m.Replace("e", 'e', 0)
	if val := m.GetOrDefault("e", 1); val != 0 {
		t.Errorf("Replace expect 0, but %v", val)
	}

	m.Replace("e", 1, 1) // won't replace
	if val := m.Replace("e", 1, 1); val != false {
		t.Errorf("Replace expect false, but true")
	}
}
