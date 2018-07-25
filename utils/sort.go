package utils

import "sort"

/*
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
*/
func Sort(objects []interface{}) {
	if objects == nil || len(objects) <= 1 {
		return
	}
	sort.Sort(sortable_struct{objects, GetComparator(objects[0])})
}

type sortable_struct struct {
	objects    []interface{}
	comparator Comparator
}

func (s sortable_struct) Len() int           { return len(s.objects) }
func (s sortable_struct) Swap(i, j int)      { s.objects[i], s.objects[j] = s.objects[j], s.objects[i] }
func (s sortable_struct) Less(i, j int) bool { return s.comparator(s.objects[i], s.objects[j]) < 0 }
