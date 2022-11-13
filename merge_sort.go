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
	r := l.next
	for r != nil {
		r = r.next
		if r == nil {
			break
		}
		r = r.next
		l = l.next
	}
	t := l.next
	l.next = nil
	for a.next != t {
		a = a.next
	}
	a.next = l
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
