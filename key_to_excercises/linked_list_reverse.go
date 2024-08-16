package main

import (
	"fmt"
	"strconv"
	"strings"
)

type LinkedListNode struct {
	Val  int
	Next *LinkedListNode
}

func ParseLinkedList(s string) *LinkedListNode {
	result := &LinkedListNode{}
	p := result

	for _, t := range strings.Split(s, "->") {
		num, err := strconv.Atoi(strings.TrimSpace(t))
		if err != nil {
			continue
		}
		p.Next = &LinkedListNode{
			Val: num,
		}
		p = p.Next
	}
	return result.Next
}

func (l *LinkedListNode) String() string {
	var nodes []string
	p := l
	for p != nil {
		nodes = append(nodes, strconv.Itoa(p.Val))
		p = p.Next
	}
	return strings.Join(nodes, "->")
}

// ReverseLinkedList 循环解法
func ReverseLinkedList(head *LinkedListNode) *LinkedListNode {
	var prev, cur *LinkedListNode
	cur = head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// ReverseLinkedListRecursive 递归解法
func ReverseLinkedListRecursive(head *LinkedListNode) *LinkedListNode {
	if head == nil || head.Next == nil {
		return head
	}
	reversed := ReverseLinkedListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return reversed
}

func ReverseLinkedListRecursiveHelper(head *LinkedListNode) (*LinkedListNode, *LinkedListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}
	reversed, tail := ReverseLinkedListRecursiveHelper(head.Next)
	head.Next = nil
	tail.Next = head
	return reversed, head
}

// ReverseLinkedListRecursive2 递归解法 not good
func ReverseLinkedListRecursive2(head *LinkedListNode) *LinkedListNode {
	reversed, _ := ReverseLinkedListRecursiveHelper(head)
	return reversed
}

func RunReverseLinkedList() {
	{
		l := ParseLinkedList("1->2->3->4->5")
		l2 := ReverseLinkedList(l)
		fmt.Println(l2)
	}

	{
		l := ParseLinkedList("1")
		l2 := ReverseLinkedList(l)
		fmt.Println(l2)
	}

	{
		l := ParseLinkedList("")
		l2 := ReverseLinkedList(l)
		fmt.Println(l2)
	}
}

func RunReverseLinkedListRecursive() {
	{
		l := ParseLinkedList("1->2->3->4->5")
		l2 := ReverseLinkedListRecursive(l)
		fmt.Println(l2)
	}

	{
		l := ParseLinkedList("1")
		l2 := ReverseLinkedListRecursive(l)
		fmt.Println(l2)
	}
	{
		l := ParseLinkedList("")
		l2 := ReverseLinkedList(l)
		fmt.Println(l2)
	}
}
