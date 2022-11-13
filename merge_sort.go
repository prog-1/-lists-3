package main

import "fmt"

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
	case b.v > a.v:
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
	*a1 = l
	for v := l; v != nil; v = v.next {
		len++
	}
	v := l
	for i := 0; i <= len/2; i++ {
		*b1, v = v, v.next
	}
	n := a
	for i := 0; i < len/2-1; i++ {
		n = n.next
	}
	n.next = nil
	return *a1, *b1
}

func mergeSort(l *node) *node {
	var i int
	for v := l; v != nil; v = v.next {
		i++
	}
	if i <= 1 {
		return l
	}
	a, b := split(l)
	return merge(mergeSort(a), mergeSort(b))
}

func main() {
	l := &node{v: 4, next: &node{v: 2, next: &node{v: 0, next: &node{v: 9}}}}
	fmt.Println(format(mergeSort(l)))
}

// func format(l *node) string {
// 	var b string
// 	for e := l; e != nil; e = e.next {
// 		b = b + fmt.Sprint(e.v)
// 		if e.next != nil {
// 			b = b + " -> "
// 		}
// 	}
// 	return b
// }
