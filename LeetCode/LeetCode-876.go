package main

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fastPoint := head
	latePoint := head

	for fastPoint.Next != nil && fastPoint.Next.Next != nil {
		fastPoint = fastPoint.Next.Next
		latePoint = latePoint.Next
	}

	if fastPoint.Next == nil {
		return latePoint
	} else {
		return latePoint.Next
	}
	return nil
}
