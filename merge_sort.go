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
	// Calculating length:
	var len int
	for i := l; i != nil; i, len = i.next, len+1 {
	}

	if len < 2 {
		panic("length must be >= 2")
	}

	a = l
	m := &l                                         // m stands for middle
	for i := 1; i < len/2; i, *m = i+1, (*m).next { // making m pointing on the last element of part a
	}
	b = (*m).next
	(*m).next = nil // Breaking connection between a and b

	return a, b
}

func mergeSort(l *node) *node {
	if l == nil || l.next == nil { // len < 1
		return l
	}
	a, b := split(l)
	a = mergeSort(a)
	b = mergeSort(b)
	return merge(a, b)
}

func main() {

}
