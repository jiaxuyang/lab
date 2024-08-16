package main

import "fmt"

// least recently used

type LRUCache struct {
	m          map[int]*Node
	head, tail *Node
	cap        int
}
type Node struct {
	Key, Val   int
	Prev, Next *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		m:    make(map[int]*Node),
		head: head,
		tail: tail,
		cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.m[key]
	if !ok {
		return -1
	}
	this.removeNode(n)
	this.makeRecent(n)
	return n.Val
}

func (this *LRUCache) removeNode(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}
func (this *LRUCache) makeRecent(n *Node) {
	n.Prev = this.tail.Prev
	n.Next = this.tail
	this.tail.Prev.Next = n
	this.tail.Prev = n
}

func (this *LRUCache) Put(key int, value int) {
	n := this.m[key]
	if n == nil {
		n = &Node{
			Key: key, Val: value,
		}
		this.m[key] = n

	} else {
		n.Val = value
		this.removeNode(n)
	}

	this.makeRecent(n)
	if len(this.m) > this.cap {
		rm := this.head.Next
		this.removeNode(rm)
		delete(this.m, rm.Key)
		// for v := this.head; v!= nil ;v= v.Next {
		//     fmt.Println(v.Key)
		// }
	}
	//fmt.Println(this.head.Next.Key)
}

func RunLru() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Get(1)
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
