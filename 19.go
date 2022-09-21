package Question_by_day

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1, p2 := head, head
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}

	if p1 == nil {
		return head.Next
	}
	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return head
}
