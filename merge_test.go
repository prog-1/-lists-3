package main

import (
	"bytes"
	"fmt"
	"testing"
)

func list(e ...int) *node {
	var l *node
	for i := len(e) - 1; i >= 0; i-- {
		l = &node{v: e[i], next: l}
	}
	return l
}

func equal(a, b *node) bool {
	for ; a != nil && b != nil; a, b = a.next, b.next {
		if a.v != b.v {
			return false
		}
	}
	return true
}

func format(l *node) string {
	var b bytes.Buffer
	for ; l != nil; l = l.next {
		fmt.Fprintf(&b, "->[%v]", l.v)
	}
	b.WriteString("->|")
	return b.String()
}

func TestMerge(t *testing.T) {
	for _, tc := range []struct {
		name               string
		init, init2, want1 *node
	}{
		{"test1", list(1, 2, 3, 4, 5, 6, 7, 8, 9), list(), list(1, 2, 3, 4, 5, 6, 7, 8, 9)},
		{"test2", list(1, 2, 3, 4, 5, 6, 7, 8, 9), list(0), list(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)},
		{"test3", list(10, 20, 30, 40), list(-10, -20, -30, -40), list(-10, -20, -30, -40, 10, 20, 30, 40)},
		{"test4", list(), list(), list()},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a, b, want := tc.init, tc.init2, tc.want1
			got := merge(a, b)
			if !equal(got, want) {
				t.Errorf("got = %v, want = %v", format(got), format(want))
			}
		})
	}
}

func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		name               string
		init, want1, want2 *node
	}{
		{"test1", list(1, 2, 3, 4, 5, 6, 7, 8, 9), list(1, 2, 3, 4), list(5, 6, 7, 8, 9)},
		{"test2", list(55, 44), list(55), list(44)},
		{"test3", list(1, 2, 54, 101, 123, 123), list(1, 2, 54), list(101, 123, 123)},
		{"test4", list(0, 1, 23, 54, 21, 21), list(0, 1, 23), list(54, 21, 21)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			l, Want, want := tc.init, tc.want1, tc.want2
			gotA, gotB := split(l)
			if !equal(gotA, Want) {
				t.Errorf("got1 = %v, Want = %v", format(gotA), format(Want))
			} else if !equal(gotB, want) {
				t.Errorf("got2 = %v, want = %v", format(gotB), format(want))
			}
		})
	}
}
func TestMergeSort(t *testing.T) {
	for _, tc := range []struct {
		name        string
		init, want1 *node
	}{
		{"test1", list(1, 2, 3, 4, 5), list(1, 2, 3, 4, 5)},
		{"test2", list(5, 4, 3, 2, 1), list(1, 2, 3, 4, 5)},
		{"test3", list(10, 6, 3, 7, 5, 2, -1, 4), list(-1, 2, 3, 4, 5, 6, 7, 10)},
		{"test4", list(100, 90, 80, 70, 60, 50, 40, 30, 20, 10), list(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, want := tc.init, tc.want1
			got = mergeSort(got)
			if !equal(got, want) {
				t.Errorf("got = %v, want = %v", format(got), format(want))
			}
		})
	}
}
