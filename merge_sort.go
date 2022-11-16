package main

import "fmt"

type node struct {
	v    int
	next *node
}

func merge(a, b *node) (c *node) {
	var d *node
	for a != nil && b != nil {
		if a.v > b.v {
			if c == nil {
				c = &node{b.v, nil}
				b = b.next
				d = c
			} else {
				d.next = &node{b.v, nil}
				b = b.next
				d = d.next
			}
		} else {
			if c == nil {
				c = &node{a.v, nil}
				a = a.next
				d = c
			} else {
				d.next = &node{a.v, nil}
				a = a.next
				d = d.next
			}
		}
	}
	if a == nil {
		d.next = b
	} else {
		d.next = a
	}
	return c
}

func split(l *node) (a, b *node) {
	a = l
	b = l
	for i := 0; l.next != nil; l, i = l.next, i+1 {
		if i%2 == 1 {
			b = b.next
		}
	}
	b.next, b = nil, b.next
	return a, b
}

func mergeSort(l *node) *node {
	if l.next == nil {
		return l
	}
	a, b := split(l)
	return merge(mergeSort(a), mergeSort(b))
}

func Print(l *node) {
	for i := l; i != nil; i = i.next {
		fmt.Println(i.v)
	}
	fmt.Println()
}

func SliceToList(s []int) (n *node) {
	for i := len(s) - 1; i >= 0; i-- {
		n = &node{s[i], n}
	}
	return n
}

func main() {
	a := mergeSort(SliceToList([]int{5, 7, 2}))
	Print(a)

}
