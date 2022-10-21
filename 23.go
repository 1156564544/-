package Question_by_day

func mergeKLists(lists []*ListNode) *ListNode {
	res := &ListNode{0, nil}
	cur := res
	for true {
		min := 100000
		idx := -1
		for i, l := range lists {
			if l == nil {
				continue
			}
			if l.Val < min {
				min = l.Val
				idx = i
			}
		}
		if min == 100000 {
			break
		}
		res.Next = &ListNode{min, nil}
		res = res.Next
		lists[idx] = lists[idx].Next
	}
	return cur.Next
}
