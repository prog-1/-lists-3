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
	}
	b.next = merge(a, b.next)
	return b
}

func split(l *node) (a, b *node) {
	var len uint
	for el := l; el != nil; el = el.next {
		len++
	}
	if len < 2 {
		return
	}
	i := len / 2
	var j uint
	for el := l; el != nil && j < len; j++ {
		if j < i {
			el = el.next
			a = el
		} else {
			el = el.next
			b = el
		}
	}
	return a, b
}

func mergeSort(l *node) *node {
	var len uint
	for el := l; el != nil; el = el.next {
		len++
	}
	if len == 1 {
		return l
	}
	a, b := split(l)
	return merge(mergeSort(a), mergeSort(b))
}
