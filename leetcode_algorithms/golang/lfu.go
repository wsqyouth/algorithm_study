package main

import "fmt"

type LinkNode struct {
	key, val, freq int
	prev, next     *LinkNode
}

type Constructor struct { //Constructor
	capacity   int
	minFreq    int
	head, tail *LinkNode
	m          map[int]*LinkNode
	freqMap    map[int]*LinkNode
}

func NewConstructor(capacity int) *Constructor {
	head := &LinkNode{-1, -1, 0, nil, nil}
	tail := &LinkNode{-1, -1, 0, nil, nil}
	head.next = tail
	tail.prev = head
	return &Constructor{capacity, 0, head, tail, make(map[int]*LinkNode), make(map[int]*LinkNode)}
}

func (this *Constructor) removeNode(node *LinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *Constructor) addNodeOnHead(node *LinkNode) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next = node
	node.next.prev = node
}

func (this *Constructor) moveNodeToHead(node *LinkNode) {
	this.removeNode(node)
	this.addNodeOnHead(node)
}

func (this *Constructor) Get(key int) int {
	node, exist := this.m[key]
	if !exist {
		return -1
	}
	this.moveNodeToHead(node)
	node.freq++
	return node.val
}

func (this *Constructor) Set(key, val int) {
	node, exist := this.m[key]
	if exist {
		node.val = val
		this.moveNodeToHead(node)
		node.freq++
	} else {
		if len(this.m) >= this.capacity {
			// 删除队尾,注意是队尾的前一个节点
			delete(this.m, this.tail.prev.key)
			this.removeNode(this.tail.prev)
		}
		node = &LinkNode{key, val, 1, nil, nil}
		this.m[key] = node
		this.addNodeOnHead(node)
		if this.minFreq == 0 || this.minFreq > 1 {
			this.minFreq = 1
		}
	}
}

func main() {
	obj := NewConstructor(2)
	obj.Set(1, 1)
	obj.Set(2, 2)

	fmt.Println("hello", obj.Get(1))
	// obj.set(3, 3)
	// fmt.Println(obj.m)
	// fmt.Println("hello", obj.get(2))
}
