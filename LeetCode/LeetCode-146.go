package main

import "fmt"

//Node defines node of cache list
type Node struct {
	Key   int
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
		deletedKey := lru.Tail.Pre.Key
		lru.Tail.Pre.Pre.Next = lru.Tail
		lru.Tail.Pre = lru.Tail.Pre.Pre
		lru.HashCache[deletedKey] = nil
	}
	lru.addNodeToHead(key, value)
}

func (lru *LRUCache) addNodeToHead(key int, value int) {
	node := &Node{
		Key:   key,
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
	obj := Constructor(1)
	obj.Put(2, 1)
	fmt.Println(obj.Get(2))
	obj.Put(3, 2)
	fmt.Println(obj.Get(2))
	fmt.Printf("head--%d--%d--tail\n", obj.Head.Next.Value, obj.Head.Next.Next.Value)
	fmt.Printf("tail--%d--%d--head\n", obj.Tail.Pre.Value, obj.Tail.Pre.Pre.Value)
	fmt.Println(obj.Get(3))
	//fmt.Println(obj.Get(4))
}
