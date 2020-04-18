package main

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preNode := &ListNode{
		Val:  0,
		Next: head,
	}
	fast := head
	slow := preNode
	count := 0
	for fast.Next != nil {
		fast = fast.Next
		count++
		if count >= n+1 {
			slow = slow.Next
		}
	}
	if count < n {
		return head
	}
	slow.Next = slow.Next.Next
	return preNode.Next

}
