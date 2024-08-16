package main

import "fmt"

func SortedLinkedListMergeRecursive(l, r *LinkedListNode) *LinkedListNode {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	if l.Val < r.Val {
		l.Next = SortedLinkedListMergeRecursive(l.Next, r)
		return l
	} else {
		r.Next = SortedLinkedListMergeRecursive(l, r.Next)
		return r
	}
}
func SortedLinkedListMerge(l, r *LinkedListNode) *LinkedListNode {
	virtual := &LinkedListNode{}
	p := virtual
	for l != nil && r != nil {
		if l.Val < r.Val {
			p.Next = l
			l = l.Next
		} else {
			p.Next = r
			r = r.Next
		}
		p = p.Next
	}
	if l != nil {
		p.Next = l
	} else if r != nil {
		p.Next = r
	}
	return virtual.Next
}

// SortedLinkedListMerge2 not good
func SortedLinkedListMerge2(l, r *LinkedListNode) *LinkedListNode {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	// 0
	// 1, 3, 5
	m := l
	if r.Val < l.Val {
		m = r
	}
	for l != nil && r != nil {
		var l2 *LinkedListNode
		for l != nil && r != nil && l.Val <= r.Val {
			l2 = l
			l = l.Next
		}
		if l2 != nil {
			l2.Next = r
		}
		var r2 *LinkedListNode
		for l != nil && r != nil && r.Val <= l.Val {
			r2 = r
			r = r.Next
		}
		if r2 != nil {
			r2.Next = l
		}
	}
	return m
}

func RunSortedLinkedListMerge() {
	{
		fmt.Println(SortedLinkedListMerge(ParseLinkedList("1->3->5"), ParseLinkedList("1->3->5")))
		fmt.Println(SortedLinkedListMerge(ParseLinkedList("1->3->5"), ParseLinkedList("")))
		fmt.Println(SortedLinkedListMerge(ParseLinkedList("1->3->5"), ParseLinkedList("2->4->6")))
		fmt.Println(SortedLinkedListMerge(ParseLinkedList("1->3->5"), ParseLinkedList("12->14->16")))
		fmt.Println(SortedLinkedListMerge(ParseLinkedList("11->13->15"), ParseLinkedList("2->4->6")))
	}
}

func RunSortedLinkedListMergeRecursive() {
	{
		fmt.Println(SortedLinkedListMergeRecursive(ParseLinkedList("1->3->5"), ParseLinkedList("1->3->5")))
		fmt.Println(SortedLinkedListMergeRecursive(ParseLinkedList("1->3->5"), ParseLinkedList("")))
		fmt.Println(SortedLinkedListMergeRecursive(ParseLinkedList("1->3->5"), ParseLinkedList("2->4->6")))
		fmt.Println(SortedLinkedListMergeRecursive(ParseLinkedList("1->3->5"), ParseLinkedList("12->14->16")))
		fmt.Println(SortedLinkedListMergeRecursive(ParseLinkedList("11->13->15"), ParseLinkedList("2->4->6")))
	}
}
