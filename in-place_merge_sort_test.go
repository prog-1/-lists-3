package main

import (
	"reflect"
	"testing"
)

func sliceToList(ints ...int) (el *node) {
	for i := len(ints) - 1; i >= 0; i-- {
		el = &node{val: ints[i], next: el}
	}
	return el
}

func listToSlice(el *node) []int {
	if el == nil {
		return nil
	}
	var s []int
	s = append(s, int(el.val))
	el = el.next
	for el != nil {
		s = append(s, int(el.val))
		el = el.next
	}
	return s
}

func TestMerge(t *testing.T) {
	for _, tc := range []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{"1", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"2", []int{3, 6, 9}, []int{1, 2, 7, 10, 12}, []int{1, 2, 3, 6, 7, 9, 10, 12}},
		{"3", []int{1}, []int{2}, []int{1, 2}},
		{"4", []int{4}, []int{4}, []int{4, 4}},
		{"5", []int{1}, []int{}, []int{1}},
		{"5", []int{}, []int{}, nil},
		{"6", nil, nil, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := listToSlice(merge(sliceToList(tc.a...), sliceToList(tc.b...)))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		a    []int
		b    []int
	}{
		{"even", []int{1, 2, 3, 6, 7, 9, 10, 12}, []int{1, 2, 3, 6}, []int{7, 9, 10, 12}},
		{"odd", []int{1, 2, 3, 6, 7, 9, 10}, []int{1, 2, 3}, []int{6, 7, 9, 10}},
		{"2 el", []int{1, 2}, []int{1}, []int{2}},
		{"same el", []int{4, 4}, []int{4}, []int{4}},
		{"one el", []int{1}, []int{1}, []int{}},
		{"empty", []int{}, nil, nil},
		{"nil", nil, nil, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a, b := split(sliceToList(tc.s...))
			if !reflect.DeepEqual(listToSlice(a), tc.a) {
				t.Errorf("got a = %v, want a = %v", a, tc.a)
			} else if !reflect.DeepEqual(listToSlice(a), tc.a) {
				t.Errorf("got b = %v, want b = %v", b, tc.b)
			}
		})
	}

}

func TestMergeSort(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want []int
	}{
		{"1", []int{6, 5, 2, 3, 1, 4}, []int{1, 2, 3, 4, 5, 6}},
		{"2", []int{7, 3, 9, 6, 4, 7, 3, -2, 3, -10}, []int{-10, -2, 3, 3, 3, 4, 6, 7, 7, 9}},
		{"3", []int{2, 1}, []int{1, 2}},
		{"4", []int{4}, []int{4}},
		{"5", []int{}, nil},
		{"6", nil, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := listToSlice(mergeSort(sliceToList(tc.s...)))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
