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
		name                  string
		initAF, initBF, wantF *node
	}{
		{"test-1", list(1, 2, 3, 4, 5, 6, 7, 8, 9), list(), list(1, 2, 3, 4, 5, 6, 7, 8, 9)},
		{"test-2", list(1, 2, 12, 32, 21, 321), list(1), list(1, 1, 2, 12, 32, 21, 321)},
		{"test-3", list(-123, 1, 2, 54, 101, 123), list(1, 2, 12, 12, 54), list(-123, 1, 1, 2, 2, 12, 12, 54, 54, 101, 123)},
		{"test-4", list(), list(), list()},
		{"test-5", list(1), list(1, 2, 12, 32, 21, 321), list(1, 1, 2, 12, 32, 21, 321)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a, b, want := tc.initAF, tc.initBF, tc.wantF
			got := merge(a, b)
			if !equal(got, want) {
				t.Errorf("got = %v, want = %v", format(got), format(want))
			}
		})
	}
}

func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		name                  string
		initF, wantAF, wantBF *node
	}{
		{"test-1", list(1, 2, 3, 4, 5, 6, 7, 8, 9), list(1, 2, 3, 4), list(5, 6, 7, 8, 9)},
		{"test-2", list(1, 2), list(1), list(2)},
		{"test-3", list(1, 2, 54, 101, 123, 123), list(1, 2, 54), list(101, 123, 123)},
		{"test-4", list(0, 1, 23, 54, 21, 21), list(0, 1, 23), list(54, 21, 21)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			l, wantA, wantB := tc.initF, tc.wantAF, tc.wantBF
			gotA, gotB := split(l)
			if !equal(gotA, wantA) {
				t.Errorf("gotA = %v, wantA = %v", format(gotA), format(wantA))
			} else if !equal(gotB, wantB) {
				t.Errorf("gotB = %v, wantB = %v", format(gotB), format(wantB))
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tc := range []struct {
		name         string
		initF, wantF *node
	}{
		{"test-1", list(1, 2, 3, 4, 5, 6, 7, 8, 9), list(1, 2, 3, 4, 5, 6, 7, 8, 9)},
		{"test-2", list(5, 4, 3, 2, 1), list(1, 2, 3, 4, 5)},
		{"test-3", list(123, 2, 1, 54, -123, 101), list(-123, 1, 2, 54, 101, 123)},
		{"test-4", list(10, 9, 8, 7, 6, 5, 4, 3, 2, 1), list(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)},
		{"test-5", list(1), list(1)},
		{"test-6", list(2, 1), list(1, 2)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, want := tc.initF, tc.wantF
			got = mergeSort(got)
			if !equal(got, want) {
				t.Errorf("got = %v, want = %v", format(got), format(want))
			}
		})
	}
}
