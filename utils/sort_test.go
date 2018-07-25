package utils

import (
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	a := []interface{}{6, 5, 3, 4, 1, 2}
	Sort(a)
	for i := 1; i < len(a); i++ {
		first := a[i-1]
		second := a[i]
		if first.(int) > second.(int) {
			t.Error("not sorted by int")
		}
	}

	a = []interface{}{'z', 'd', 'g', 'a'}
	Sort(a)
	for i := 1; i < len(a); i++ {
		first := a[i-1]
		second := a[i]
		if first.(rune) > second.(rune) {
			t.Error("not sorted by rune")
		}
	}

	a = []interface{}{"z", "d", "g", "a"}
	Sort(a)
	for i := 1; i < len(a); i++ {
		first := a[i-1]
		second := a[i]
		if first.(string) > second.(string) {
			t.Error("not sorted by string")
		}
	}

	now := time.Now()
	a = []interface{}{now.Add(10), now.Add(-10), now}
	Sort(a)
	for i := 1; i < len(a); i++ {
		first := a[i-1]
		second := a[i]
		if first.(time.Time).After(second.(time.Time)) {
			t.Error("not sorted by time.Time")
		}
	}

	a = []interface{}{1.6, -1.0, 9.99}
	Sort(a)
	for i := 1; i < len(a); i++ {
		first := a[i-1]
		second := a[i]
		switch first.(type) {
		case float32:
			if first.(float32) > second.(float32) {
				t.Error("not sorted by float32")
			}
		case float64:
			if first.(float64) > second.(float64) {
				t.Error("not sorted by float64")
			}
		}

	}

}
