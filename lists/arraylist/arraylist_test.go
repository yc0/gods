package arraylist

import (
	"testing"
)

func TestListAdd(t *testing.T) {
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

	list.AddAt(2, 10)
	if actualValue, _ := list.Get(3); actualValue != 9 {
		t.Errorf("Got %v expected %v", actualValue, 9)
	}
	list.Add(11, 12, 13, 14, 15, 16)
	list.AddAt(9, 17)
	if actualValue, _ := list.Get(list.Size() - 1); actualValue != 16 {
		t.Errorf("Got %v expected %v", actualValue, 16)
	}
}

func TestListContains(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(2)
	if actualValue := list.Contains(2); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Contains(nil); actualValue != true {
		t.Errorf("%+v", list)
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Contains(3); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.IndexOf(2); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

}

func TestListClone(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(2)

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

func TestListRemove(t *testing.T) {
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
	list.Remove(3)
	t.Logf("%+v", list)
}

func TestListSort(t *testing.T) {
	list := New()
	list.Add(6, 8, 3, 9, 4, 5, 2, 1, 7)
	list.Add(-1)
	list.Sort()
	for i := 1; i < list.Size(); i++ {
		first, _ := list.Get(i - 1)
		second, _ := list.Get(i)
		if first.(int) > second.(int) {
			t.Error("not sorted by int")
		}
	}
	t.Logf("%+v", list)
}

func BenchmarkListInsert(b *testing.B) {
	l := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Insert(0, i)
	}
}
func BenchmarkListAddAt(b *testing.B) {
	l := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.AddAt(0, i)
	}
}
