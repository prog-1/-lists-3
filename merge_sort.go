package main

import (
	"fmt"
)

type node struct {
	v    int
	next *node
}

func merge(a, b *node) *node {
	switch {
	case a == nil:
		return b
	case b == nil:
		return a
	case a.v < b.v:
		a.next = merge(a.next, b)
		return a
	default:
		b.next = merge(a, b.next)
		return b
	}
}

func split(l *node) (a, b *node) {
	var len int
	a1, b1 := &a, &b
	for n := l; n != nil; n = n.next {
		len++
	}
	n := l
	for i := 0; i <= len/2 && l.next != nil; i++ {
		*b1, n = n, n.next
	}

	*a1 = l
	n = a
	for i := 0; i < len/2-1 && l.next != nil; i++ {
		n = n.next
	}
	n.next = nil
	return a, b
}

func mergeSort(l *node) *node {
	if l.next == nil {
		return l
	}
	a, b := split(l)
	l = merge(mergeSort(a), mergeSort(b))
	return l
}

func print(a *node) {
	for ; a != nil; a = a.next {
		fmt.Print(" --> ", a.v)
	}
	fmt.Println(" --> |")
}

func main() {
	l := &node{2, &node{4, &node{6, &node{3, &node{1, &node{8, nil}}}}}}
	a1 := mergeSort(l)
	print(a1)

}
