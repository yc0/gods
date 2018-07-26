package vector

import (
	"testing"
)

func TestVectorAdd(t *testing.T) {
	vector := New()
	if actualValue := vector.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	vector.Add(1)
	vector.Add(7)
	vector.Add(9)
	if actualValue := vector.IsEmpty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := vector.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	if _, err := vector.Get(3); err == nil {
		t.Errorf("Expect exception %v", err.Error())
	}

	vector.AddAt(2, 10)
	if actualValue, _ := vector.Get(3); actualValue != 9 {
		t.Errorf("Got %v expected %v", actualValue, 9)
	}
	vector.Add(11, 12, 13, 14, 15, 16)
	vector.AddAt(9, 17)
	if actualValue, _ := vector.Get(vector.Size() - 1); actualValue != 16 {
		t.Errorf("Got %v expected %v", actualValue, 16)
	}
}

func TestVectorContains(t *testing.T) {
	vector := New()
	vector.Add(1)
	vector.Add(2)
	if actualValue := vector.Contains(2); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := vector.Contains(nil); actualValue != true {
		t.Errorf("%+v", vector)
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := vector.Contains(3); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := vector.IndexOf(2); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

}

func TestVectorClone(t *testing.T) {
	vector := New()
	vector.Add(1)
	vector.Add(2)

	second := vector.Clone()
	if vector == second {
		t.Errorf("Got two vector is the same :%v expected %v, %p and %p", true, false, vector, second)
	}
	if actualValue := second.Contains(2); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := second.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
}

func TestVectorRemove(t *testing.T) {
	vector := New()
	vector.Add(1)
	vector.Add(2)

	if actualValue, _ := vector.Remove(1); actualValue != 2 {
		t.Errorf("Remove %v expected %v", actualValue, 2)
	}

	if nil, err := vector.Remove(1); err == nil {
		t.Errorf("Remove expected %v", err.Error())
	}
	vector.Add(3, 4, 5, 6)
	vector.Remove(3)
	t.Logf("%+v", vector)
}

func TestVectorSort(t *testing.T) {
	vector := New()
	vector.Add(6, 8, 3, 9, 4, 5, 2, 1, 7)
	vector.Add(-1)
	vector.Sort()
	for i := 1; i < vector.Size(); i++ {
		first, _ := vector.Get(i - 1)
		second, _ := vector.Get(i)
		if first.(int) > second.(int) {
			t.Error("not sorted by int")
		}
	}
	t.Logf("%+v", vector)
}

func TestVectorMutexOperate(t *testing.T) {

}
