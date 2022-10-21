package Question_by_day

func numComponents(head *ListNode, nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	res := 0
	flag := false
	for head != nil {
		if m[head.Val] {
			if flag == false {
				res += 1
				flag = true
			}
		} else {
			flag = false
		}
		head = head.Next
	}

	return res
}
