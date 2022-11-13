package main

type node struct {
	v    int
	next *node
}

func merge(a, b *node) *node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.v < b.v {
		a.next = merge(a.next, b)
		return a
	} else {
		b.next = merge(a, b.next)
		return b
	}
}

func split(l *node) (a, b *node) {
	len := 0
	for n := l; n != nil; n = n.next {
		len++
	}
	a = l
	xd := &l
	for e := 1; e < len/2; e, *xd = e+1, (*xd).next {
	}
	b = (*xd).next
	(*xd).next = nil

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

func main() {
}
