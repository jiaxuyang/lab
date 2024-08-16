package main

type LFUCache struct {
	cap, size int
	minFreq   int
	keyMap    map[int]*LFUNode
	freqMap   map[int]*LFUNodeList
}

type LFUNode struct {
	key, value, freq int
	prev, next       *LFUNode
}

type LFUNodeList struct {
	head, tail *LFUNode
}

func NewLFUNodeList() *LFUNodeList {
	head, tail := &LFUNode{}, &LFUNode{}
	head.next = tail
	tail.prev = head
	return &LFUNodeList{
		head: head,
		tail: tail,
	}
}

func LFUConstructor(capacity int) LFUCache {
	return LFUCache{
		cap:     capacity,
		size:    0,
		minFreq: 0,
		keyMap:  make(map[int]*LFUNode),
		freqMap: make(map[int]*LFUNodeList),
	}
}

func (this *LFUCache) removeNode(n *LFUNode) {
	n.prev.next = n.next
	n.next.prev = n.prev
}
func (this *LFUCache) appendNode(n *LFUNode) {
	if this.freqMap[n.freq] == nil {
		this.freqMap[n.freq] = NewLFUNodeList()
	}
	l := this.freqMap[n.freq]
	n.prev = l.tail.prev
	n.next = l.tail
	l.tail.prev.next = n
	l.tail.prev = n
}

func (this *LFUCache) Get(key int) int {
	n := this.keyMap[key]
	if n == nil {
		return -1
	}
	this.removeNode(n)
	if nodeList := this.freqMap[n.freq]; nodeList.head.next == nodeList.tail && this.minFreq == n.freq {
		this.minFreq++
	}
	n.freq++
	this.appendNode(n)

	return n.value
}

func (this *LFUCache) Put(key int, value int) {
	n := this.keyMap[key]
	if n == nil {
		// 添加之前删除
		if this.size == this.cap {
			toDel := this.freqMap[this.minFreq].head.next
			if toDel != nil {
				this.removeNode(toDel)
				delete(this.keyMap, toDel.key)
				this.size--
			}
		}
		n = &LFUNode{
			key:   key,
			value: value,
			freq:  1,
			prev:  nil,
			next:  nil,
		}
		this.minFreq = n.freq // 新加的肯定是访问频率最低的了
		this.appendNode(n)
		this.keyMap[n.key] = n
	} else {
		this.removeNode(n)
		if nodeList := this.freqMap[n.freq]; nodeList.head.next == nodeList.tail && this.minFreq == n.freq {
			this.minFreq++
		}
		n.value = value
		n.freq++
		this.appendNode(n)
	}
}
