package main

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	for _, tc := range []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 3, 2}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 5, 3, 4, 0}, []int{0, 1, 3, 4, 5}},
	} {
		input1 := createListFromArray(tc.input)
		want1 := createListFromArray(tc.want)

		got := mergeSort(input1)
		if !equal(want1, got) {
			t.Errorf("input %v; want %v ; got %v", tc.input, tc.want, listToSlice(got))
		}
	}
}

func listToSlice(l *node) []int {
	s := make([]int, 0)
	for ; l != nil; l = l.next {
		s = append(s, l.v)
	}
	return s
}

func createListFromArray(s []int) (l *node) {
	a := &l
	for i := range s {
		*a = &node{s[i], nil}
		a = &(*a).next
	}
	return l
}

func equal(a, b *node) bool {
	for ; a != nil && b != nil; a, b = a.next, b.next {
		if a.v != b.v {
			return false
		}
	}
	return a == b
}
