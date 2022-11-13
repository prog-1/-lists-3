package main

import (
	"testing"
)

func list1(e ...int) *node {
	var l *node
	for i := len(e) - 1; i >= 0; i-- {
		l = &node{v: e[i]}
	}
	return l
}

func TestMerge(t *testing.T) {
	for _, tc := range []struct {
		name string
		a, b *node
		want *node
	}{
		{"one of the lists is empty", list1(1, 2, 7, 8, 9), list1(), list1(1, 2, 7, 8, 9)},
		{"basic", list1(1, 2, 3), list1(4, 5, 6), list1(1, 2, 3, 4, 5, 6)},
		{"numbers repeat", list1(1, 3, 9, 9), list1(9, 11), list1(1, 3, 9, 9, 9, 11)},
		{"merge", list1(1, 3, 9), list1(4, 5), list1(1, 3, 4, 5, 9)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a, b, want := tc.a, tc.b, tc.want
			got := merge(a, b)
			if !equal(got, want) {
				t.Errorf("got = %v, want = %v", got, want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		name         string
		a            *node
		want1, want2 *node
	}{
		{"two numbers", list1(1, 1), list1(1), list1(1)},
		{"negative numbers", list1(9, 7, -3, 4, 0), list1(9, 7), list1(-3, 4, 0)},
		{"another case", list1(133, 1, 1, 3, 4), list1(133, 1), list1(1, 3, 4)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a, want1, want2 := tc.a, tc.want1, tc.want2
			got1, got2 := split(a)
			if !equal(got1, want1) {
				t.Errorf("got = %v, want = %v", got1, want1)
			} else if !equal(got2, want2) {
				t.Errorf("got = %v, want = %v", got2, want1)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tc := range []struct {
		name string
		a    *node
		want *node
	}{
		{"sorted", list1(1, 2, 3, 4, 6), list1(1, 2, 3, 4, 6)},
		{"unsorted", list1(1, 9, 6, 3, 2), list1(1, 2, 3, 6, 9)},
		{"unsorted, numbers repeat", list1(1, 9, 9, 6, 1, 6, 2, 1), list1(1, 1, 1, 2, 6, 9, 9)},
		//have problems with this test
		{"negative numbers", list1(-4, 1, 5, 3, -5), list1(-5, -4, 1, 3, 5)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a, want := tc.a, tc.want
			got := mergeSort(a)
			if !equal(got, want) {
				t.Errorf("got = %v, want = %v", got, want)
			}
		})
	}
}

func equal(a, b *node) bool {
	for ; a != nil && b != nil; a, b = a.next, b.next {
		if a.v != b.v {
			return false
		}
	}
	return true
}
