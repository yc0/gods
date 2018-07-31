package hashmap

import (
	"testing"
)

func TestNew(t *testing.T) {
	m := New()
	m.m['a'] = 1
	t.Errorf("%+v", m)
}
