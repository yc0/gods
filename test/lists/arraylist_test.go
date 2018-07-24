package arraylist

import (
	"testing"

	. "github.com/yc0/gods/lists/arraylist"
)

func TestListAdd(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(7)
	list.Add(9)
}
