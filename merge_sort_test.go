package main

import (
	"bytes"
	"fmt"
	"testing"
)

func list(e ...int) (l *node) {
	for i := len(e) - 1; i >= 0; i-- {
		l = &node{v: e[i], next: l}
	}
	return l
}

func equal(a, b *node) bool {
	for ; a != nil; a, b = a.next, b.next {
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
		{"case-1", list(1, 2), list(3, 4), list(1, 2, 3, 4)},
		{"case-2", list(1, 4), list(2, 3, 5), list(1, 2, 3, 4, 5)},
		{"case-3", list(1, 3), list(2, 3, 4), list(1, 2, 3, 3, 4)},
		{"case-4", list(1, 2, 3, 4, 5), list(), list(1, 2, 3, 4, 5)},
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
		{"case-1", list(2, 1, 3, 4), list(2, 1), list(3, 4)},
		{"case-2", list(2, 1, 3, 4, 0), list(2, 1), list(3, 4, 0)},
		{"case-3", list(2, 1, 1, 3, 4), list(2, 1), list(1, 3, 4)},
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
		{"case-1", list(2, 1, -1), list(-1, 1, 2)},
		{"case-2", list(1, 2, 0, 5, 4, 3), list(0, 1, 2, 3, 4, 5)},
		{"case-3", list(-1, 2, 0, 5, 4, 3, 6), list(-1, 0, 2, 3, 4, 5, 6)},
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
