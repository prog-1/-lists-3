package main

import (
	"fmt"
	"testing"
)

func initList(e ...int) *node {
	if len(e) == 0 {
		return nil
	}
	return &node{v: e[0], next: initList(e[1:]...)}
}

func equal(a, b *node) bool {
	var al, bl int
	for v := a; v != nil; v = v.next {
		al++
	}
	for v := b; v != nil; v = v.next {
		bl++
	}
	if al != bl {
		return false
	}
	for ; a != nil; a, b = a.next, b.next {
		if a.v != b.v {
			return false
		}
	}
	return true
}

func TestMerge(t *testing.T) {
	for _, tc := range []struct {
		initAF, initBF, wantF *node
	}{
		{initList(), initList(), initList()},
		{initList(1), initList(3), initList(1, 3)},
		{initList(-2, -3, 0), initList(1, 11, -20), initList(-20, -3, -2, 0, 1, 11)},
		{initList(1, 2), initList(3, 4), initList(1, 2, 3, 4)},
		{initList(1, 4), initList(2, 3, 5), initList(1, 2, 3, 4, 5)},
		{initList(1, 3), initList(2, 3, 4), initList(1, 2, 3, 3, 4)},
	} {
		a, b, want := tc.initAF, tc.initBF, tc.wantF
		got := merge(a, b)
		if !equal(got, want) {
			t.Errorf("got = %v, want = %v", format(got), format(want))
		}
	}
}

func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		initF, wantAF, wantBF *node
	}{
		{initList(1, 1), initList(1), initList(1)},
		{initList(2, -3, 0, 99), initList(2, -3), initList(0, 99)},
		{initList(2, 1, 3, 4), initList(2, 1), initList(3, 4)},
		{initList(2, 8, 3, 4, 0), initList(2, 8), initList(3, 4, 0)},
		{initList(2, 12, 1, 3, 4), initList(2, 12), initList(1, 3, 4)},
	} {
		l, wantA, wantB := tc.initF, tc.wantAF, tc.wantBF
		gotA, gotB := split(l)
		if !equal(gotA, wantA) {
			t.Errorf("gotA = %v, wantA = %v", format(gotA), format(wantA))
		} else if !equal(gotB, wantB) {
			t.Errorf("gotB = %v, wantB = %v", format(gotB), format(wantB))
		}

	}
}

func TestMergeSort(t *testing.T) {
	for _, tc := range []struct {
		initF, wantF *node
	}{
		{initList(), initList()},
		{initList(1), initList(1)},
		{initList(1000, -7), initList(-7, 1000)},
		{initList(2, 1, -1), initList(-1, 1, 2)},
		{initList(1, 2, 0, 5, 4, 3), initList(0, 1, 2, 3, 4, 5)},
		{initList(-1, 2, 0, 5, 4, 3, 6), initList(-1, 0, 2, 3, 4, 5, 6)},
	} {
		got, want := tc.initF, tc.wantF
		got = mergeSort(got)
		if !equal(got, want) {
			t.Errorf("got = %v, want = %v", format(got), format(want))
		}
	}
}

func format(l *node) string {
	var b string
	for e := l; e != nil; e = e.next {
		b = b + fmt.Sprint(e.v)
		if e.next != nil {
			b = b + " -> "
		}
	}
	return b
}
