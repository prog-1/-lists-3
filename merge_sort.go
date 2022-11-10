package main

import "fmt"

type node struct {
	v    int
	next *node
}

func SliceToList(s []int) (n *node) {
	for i := len(s) - 1; i >= 0; i-- {
		n = &node{s[i], n}
	}
	return n
}
func ListToSlice(l *node) []int {
	s := make([]int, 0)
	for cur := l; cur != nil; cur = cur.next {
		s = append(s, cur.v)
	}
	return s
}
func Print(l *node) {
	for i := l; i != nil; i = i.next {
		fmt.Println(i.v)
	}
	fmt.Println()
}

// How to merge?
// 1. Take the smallest element out of a and b
// 2.
func merge(a, b *node) (head *node) {
	for cur := &head; a != nil || b != nil; cur = &(*cur).next {
		if a == nil {
			*cur = b
			b = b.next
		} else if b == nil {
			*cur = a
			a = a.next
		} else if a.v < b.v {
			*cur = a
			a = a.next
		} else {
			*cur = b
			b = b.next
		}
	}

	return head
}

func split(l *node) (a, b *node) {
	// TODO
	fmt.Println("Smth")
	return nil, nil
}

func mergeSort(l *node) *node {
	// TODO
	return nil
}

func main() {
	a, b := SliceToList([]int{1}), SliceToList([]int{2})
	l := merge(a, b)
	Print(l)
}
