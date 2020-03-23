package main

//ListNode used to list
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	newHead := &ListNode{
		Val:  0,
		Next: head,
	}

	fastPoint := newHead.Next
	latePoint := newhead

	for fastPoint.Next != nil {
		if fastPoint.Next != nil && fastPoint.Val == fastPoint.Next.Val {
			for fastPoint.Val == fastPoint.Next.Val && fastPoint.Next != nil {
				fastPoint = fastPoint.Next
			}
			fastPoint = fastPoint.Next
			latePoint.Next = fastPoint

			if fastPoint == nil {
				break
			}
		} else {
			latePoint = latePoint.Next
			fastPoint = fastPoint.Next
		}
	}

	return newHead.Next
}
