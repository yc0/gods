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
	// if actualValue := second.Contains(2); actualValue != true {
	// 	t.Errorf("Got %v expected %v", actualValue, true)
	// }
	// if actualValue := second.Size(); actualValue != 2 {
	// 	t.Errorf("Got %v expected %v", actualValue, 2)
	// }
}

func TestListRemove(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(2)
	t.Logf("%+v", list)
}
