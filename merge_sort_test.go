package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type Want struct {
		a, b []int
	}
	for _, tc := range []struct {
		name string
		init []int
		want Want
	}{
		{"empty", []int{}, Want{[]int{}, []int{}}},
		{"1", []int{1, 2}, Want{[]int{1}, []int{2}}},
	} {
		a, b := split(SliceToList(tc.init))
		if s := ListToSlice(a); !reflect.DeepEqual(s, tc.want.a) {
			t.Errorf("want = %v, got = %v", tc.want.a, s)
		}
		if s := ListToSlice(b); !reflect.DeepEqual(s, tc.want.b) {
			t.Errorf("want = %v, got = %v", tc.want.b, s)
		}
	}
}

func TestMerge(t *testing.T) {
	type Input struct {
		a, b []int
	}
	for _, tc := range []struct {
		name  string
		input Input
		want  []int
	}{
		{"empty", Input{[]int{}, []int{}}, []int{}},
		{"1", Input{[]int{1}, []int{2}}, []int{1, 2}},
		{"2", Input{[]int{1, 4}, []int{2, 3}}, []int{1, 2, 3, 4}},
	} {
		got := merge(SliceToList(tc.input.a), SliceToList(tc.input.b))
		if s := ListToSlice(got); !reflect.DeepEqual(s, tc.want) {
			t.Errorf("want = %v, got = %v", tc.want, s)
		}
	}
}

func TestMergeSort(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []int
		want  []int
	}{
		{"empty", []int{}, []int{}},
		{"nil", nil, nil},
		{"1", []int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
	} {
		got := mergeSort(SliceToList(tc.input))
		if s := ListToSlice(got); !reflect.DeepEqual(s, tc.want) {
			t.Errorf("got = %v, want = %v", s, tc.want)
		}
	}
}
