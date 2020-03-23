package main

import (
	"fmt"
	"math"
	"sort"
)

type Node struct {
	Value int
	Pre   *Node
	Next  *Node
}

type MaxStack struct {
	Max    []int
	Buttom *Node
	Head   *Node
}

/** initialize your data structure here. */
func Constructor() MaxStack {
	this := MaxStack{
		Max: []int{},
		Buttom: &Node{
			Value: math.MinInt64,
			Pre:   nil,
			Next:  nil,
		},
		Head: &Node{
			Value: math.MinInt64,
			Pre:   nil,
			Next:  nil,
		},
	}
	this.Head.Pre = this.Buttom
	return this
}

func (this *MaxStack) Push(x int) {
	tmp := &Node{
		Value: x,
		Pre:   this.Head.Pre,
		Next:  this.Head,
	}
	this.Head.Pre.Next = tmp
	this.Head.Pre = tmp
	this.Max = append(this.Max, x)
	sort.Ints(this.Max)
}

func (this *MaxStack) Pop() int {
	res := this.Head.Pre.Value
	length := len(this.Max)
	if this.Head.Pre != this.Buttom {
		this.Head.Pre.Pre.Next = this.Head
		this.Head.Pre = this.Head.Pre.Pre
	}
	if res == this.Max[length-1] {
		this.Max = this.Max[:length-1]
	}
	if res == 54 {
		fmt.Print(this.Max[len(this.Max)-1])
	}

	return res
}

func (this *MaxStack) Top() int {
	return this.Head.Pre.Value
}

func (this *MaxStack) PeekMax() int {
	length := len(this.Max)
	if length > 0 {
		return this.Max[len(this.Max)-1]
	}
	return math.MinInt64
}

func (this *MaxStack) PopMax() int {
	tmp := this.Head.Pre
	length := len(this.Max)
	maxValue := this.PeekMax()
	for tmp != this.Buttom {
		if tmp.Value == maxValue {
			tmp.Next.Pre = tmp.Pre
			tmp.Pre.Next = tmp.Next
			break
		}
		tmp = tmp.Pre
	}
	this.Max = this.Max[:length-1]
	return maxValue
}

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */
