package linkedlist

import (
	"testing"
)

func TestLinkedListAppend(t *testing.T) {
	list := New()
	if actualValue := list.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	list.Add(1)
	list.Add(7)
	list.Add(9)
	if actualValue := list.IsEmpty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if _, err := list.Get(3); err == nil {
		t.Errorf("Expect exception %v", err.Error())
	}

	if v, _ := list.Get(2); v != 9 {
		t.Errorf("Got %v expected %v", v, 9)
	}

	list.AddAt(2, 10)
	if actualValue, _ := list.Get(3); actualValue != 9 {
		t.Errorf("Got %v", list)
		v, _ := list.Get(3)
		t.Errorf("Got %v", v)
		t.Errorf("Got %v expected %v", actualValue, 9)
	}
	list.Add(11, 12, 13, 14, 15, 16)
	list.AddAt(9, 17)
	if actualValue, _ := list.Get(list.Size() - 1); actualValue != 16 {
		t.Errorf("Got %v expected %v", actualValue, 16)
	}
}

func TestLinkedListContains(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(2)
	if actualValue := list.Contains(2); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	if actualValue := list.Contains(3); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.IndexOf(2); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}
func TestLinkedListClone(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(2)

	if actualValue := list.Contains(2); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	second := list.Clone()
	if list == second {
		t.Errorf("Got two list is the same :%v expected %v, %p and %p", true, false, list, second)
	}
	if actualValue := second.Contains(2); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := second.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
}
func TestLinkedListRemove(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(2)

	if actualValue, _ := list.Remove(1); actualValue != 2 {
		t.Errorf("Remove %v expected %v", actualValue, 2)
	}

	if nil, err := list.Remove(1); err == nil {
		t.Errorf("Remove expected %v", err.Error())
	}
	list.Add(3, 4, 5, 6)
	if actualValue, _ := list.Remove(3); actualValue != 5 {
		t.Errorf("Remove %v expected %v", actualValue, 5)
	}
}

func TestLinkedListSort(t *testing.T) {
	list := New()
	list.Add(7, 3, 0, 2, -1, 4, 9, 5)
	list.Sort()
	for cur := list.first.next; cur == list.last; cur = cur.next {
		if cur.data.(int) < cur.previous.data.(int) {
			t.Errorf("Not sorted")
		}
	}
}
