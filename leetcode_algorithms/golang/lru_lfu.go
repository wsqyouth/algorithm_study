package main

import (
	"fmt"
)

type LinkNode struct {
	key, val   int
	prev, next *LinkNode
}
type LRUCache struct {
	m          map[int]*LinkNode //存储链表的指针,省去了在链表查找节点前后地址的步骤
	capacity   int
	head, tail *LinkNode //首尾哨兵节点简化了链表的插入和删除操作
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{-1, -1, nil, nil}
	tail := &LinkNode{-1, -1, nil, nil}
	head.next = tail
	tail.prev = head
	return LRUCache{m: make(map[int]*LinkNode), capacity: capacity, head: head, tail: tail}
}

func (this *LRUCache) addNodeOnHead(node *LinkNode) {
	node.prev = this.head
	node.next = this.head.next //头插保持最新
	this.head.next = node
	node.next.prev = node
}
func (this *LRUCache) removeNode(node *LinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (this *LRUCache) moveNodeToHead(node *LinkNode) {
	this.removeNode(node)
	this.addNodeOnHead(node)
}

// Get 获取时若存在则将其调整为最新
func (this *LRUCache) Get(key int) int {
	m := this.m
	if node, ok := m[key]; ok {
		this.moveNodeToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	m := this.m
	if node, ok := m[key]; ok {
		node.val = value
		this.moveNodeToHead(node)
	} else {
		n := &LinkNode{key: key, val: value, prev: nil, next: nil}
		if len(m) >= this.capacity {
			delete(m, this.tail.prev.key) //小心
			this.removeNode(this.tail.prev)
		}
		m[key] = n
		this.addNodeOnHead(n)
	}
}

func main() {
	/**
	 * Your LRUCache object will be instantiated and called as such:
	 */
	obj := Constructor(1)
	obj.Put(2, 1)
	fmt.Println(obj.Get(2))
	obj.Put(3, 2)
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
}

/*
我们使用哈希表来存储「键值对」，键值对的键作为哈希表的 Key，而哈希表的 Value 则使用我们自己封装的 Node 类，Node 同时作为双向链表的节点。

插入：检查当前键值对是否已经存在于哈希表：
如果存在，则更新键值对，并将当前键值对所对应的 Node 节点调整到链表头部（refresh 操作）
如果不存在，则检查哈希表容量是否已经达到容量：
没达到容量：插入哈希表，并将当前键值对所对应的 Node 节点调整到链表头部（refresh 操作）
已达到容量：先从链表尾部找到待删除元素进行删除（delete 操作），然后再插入哈希表，并将当前键值对所对应的 Node 节点调整到链表头部（refresh 操作）
查询：如果没在哈希表中找到该 Key，直接返回 -1−1；如果存在该 Key，则将对应的值返回，并将当前键值对所对应的 Node 节点调整到链表头部（refresh 操作）
一些细节：

为了减少双向链表左右节点的「判空」操作，我们预先建立两个「哨兵」节点 head 和 tail。

*/
