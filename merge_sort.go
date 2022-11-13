package main

type node struct {
	v    int
	next *node
}

func merge(a, b *node) *node {
	if a == nil {
		return b
	} else if b == nil {
		return a
	} else if a.v < b.v {
		a.next = merge(a.next, b)
		return a
	} else if a.v > b.v {
		b.next = merge(a, b.next)
		return b
	}
	return nil
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

func main() {
}
