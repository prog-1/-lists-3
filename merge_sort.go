package main

type node struct {
	v    int
	next *node
}

func merge(a, b *node) *node {
	l := new(node)
	cur := l
	for a != nil && b != nil {
		if a.v <= b.v {
			cur.next = a
			a = a.next
		} else {
			cur.next = b
			b = b.next
		}
		cur = cur.next
	}
	if a != nil {
		cur.next = a
	}
	if b != nil {
		cur.next = b
	}
	return l.next
}

func split(l *node) (a, b *node) {
	a, b = new(node), new(node)
	curA, curB := a, b
	var len int
	for tmp := l; tmp != nil; tmp = tmp.next {
		len++
	}
	if len < 2 {
		panic("length must be >= 2")
	}
	for i := 0; l != nil; i++ {
		if i < len/2 {
			curA.next = l
			curA = curA.next
		} else {
			curB.next = l
			curB = curB.next
		}
		l = l.next
	}
	curA.next = nil
	return a.next, b.next
}

func mergeSort(l *node) *node {
	if l.next == nil {
		return l
	}
	a, b := split(l)
	return merge(mergeSort(a), mergeSort(b))
}
