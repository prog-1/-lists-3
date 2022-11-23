package main

import "fmt"

type node struct {
	val  int
	next *node
}

func merge(a, b *node) *node {

	//start of the list
	var res *node
	//looking for first element
	if a == nil && b == nil {
		return nil
	}
	if a == nil {
		res, b = b, b.next // if a = nil then in res += b and take next b
	} else if b == nil {
		res, a = a, a.next // if b = nil then in res += a and take next a
	} else if a.val < b.val {
		res, a = a, a.next // if a < b then in res += a and take next a
	} else {
		res, b = b, b.next // if a > b then in res += b and take next b
	}

	// looping through every element in a and b
	el := res
	for a != nil || b != nil {
		if a == nil {
			el.next, b = b, b.next // if a = nil then in res += b and take next b
		} else if b == nil {
			el.next, a = a, a.next // if b = nil then in res += a and take next a
		} else if a.val < b.val {
			el.next, a = a, a.next // if a < b then in res += a and take next a
		} else {
			el.next, b = b, b.next // if a > b then in res += b and take next b
		}
		el = el.next
	}
	return res
}

func split(l *node) (a, b *node) {

	//edge case
	if l == nil {
		return nil, nil
	}

	//make a slice from start to middle
	a = l

	//get lenght
	var len int
	for ; l != nil; l = l.next {
		len++
	}

	//get middle node
	l = a
	for i := 0; i < (len/2)-1; i++ {
		if l.next != nil {
			l = l.next
		}
	}

	//make b slice from middle to end
	b = l.next

	//break connection between a and b
	l.next = nil

	return a, b
}

func mergeSort(l *node) *node {

	// base case
	if l == nil || l.next == nil {
		return l
	}

	//split
	ls, rs := split(l)

	//recursion and merge
	ls = mergeSort(ls)
	rs = mergeSort(rs)
	return merge(ls, rs)

}

func SliceToList(s []int) (el *node) {
	for i := len(s) - 1; i >= 0; i-- {
		el = &node{val: s[i], next: el}
	}
	return el
}

func PrintList(el *node) {
	for ; el != nil; el = el.next {
		fmt.Print(el.val, " ")
	}
	fmt.Println()
}

func main() {
	//a, b := split(SliceToList([]int{4, 2, 7, 1}))
	//PrintList(a)
	//PrintList(b)
	//PrintList(merge(a, b))
	PrintList(mergeSort(SliceToList([]int{})))
}
