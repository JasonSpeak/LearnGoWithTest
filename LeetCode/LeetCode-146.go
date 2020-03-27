package main

import "fmt"

//Node defines node of cache list
type Node struct {
	Value int
	Pre   *Node
	Next  *Node
}

//LRUCache defines LRU cache
type LRUCache struct {
	MaxCapacity     int
	CurrentCapacity int
	HashCache       map[int]*Node
	Head            *Node
	Tail            *Node
}

//Constructor defines the constructor for LRUCache
func Constructor(capacity int) LRUCache {
	HeadNode := &Node{
		Value: 0,
		Pre:   nil,
		Next:  nil,
	}
	TailNode := &Node{
		Value: 0,
		Pre:   HeadNode,
		Next:  nil,
	}
	HeadNode.Next = TailNode
	return LRUCache{
		MaxCapacity:     capacity,
		CurrentCapacity: 0,
		HashCache:       map[int]*Node{},
		Head:            HeadNode,
		Tail:            TailNode,
	}
}

//Get defines function to get value
func (this *LRUCache) Get(key int) int {
	keyNode := this.HashCache[key]
	if keyNode != nil {
		this.resetNode(keyNode)
		return keyNode.Value
	}
	return -1
}

//Put defines function to put value
func (this *LRUCache) Put(key int, value int) {
	keyNode := this.HashCache[key]
	if keyNode != nil {
		this.updateNodeValue(keyNode, value)
	} else {
		this.addNewNode(key, value)
	}
}

func (lru *LRUCache) resetNode(node *Node) {
	if lru.CurrentCapacity > 1 {
		node.Pre.Next = node.Next
		fmt.Println("Node next value is", node.Next.Value)
		node.Next.Pre = node.Pre
		lru.Head.Next.Pre = node
		node.Next = lru.Head.Next
		lru.Head.Next = node
		node.Pre = lru.Head
	}
}

func (lru *LRUCache) updateNodeValue(node *Node, value int) {
	node.Value = value
	lru.resetNode(node)
}

func (lru *LRUCache) addNewNode(key int, value int) {
	if lru.CurrentCapacity < lru.MaxCapacity {
		lru.CurrentCapacity++
	} else {
		deletedValue := lru.Tail.Pre.Value
		fmt.Println("DeletedValue is ", deletedValue)
		lru.Tail.Pre.Pre.Next = lru.Tail
		lru.Tail.Pre = lru.Tail.Pre.Pre
		lru.Tail.Pre.Pre = nil
		lru.Tail.Pre.Next = nil
		lru.HashCache[deletedValue] = nil
	}
	lru.addNodeToHead(key, value)
}

func (lru *LRUCache) addNodeToHead(key int, value int) {
	node := &Node{
		Value: value,
		Pre:   lru.Head,
		Next:  lru.Head.Next,
	}
	lru.Head.Next.Pre = node
	lru.Head.Next = node
	lru.HashCache[key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	//fmt.Println(obj.Get(4))
}
