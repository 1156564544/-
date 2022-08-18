package Question_by_day

func deepestLeavesSum(root *TreeNode) int {
	deque := []*TreeNode{root}
	for len(deque) > 0 {
		curSize := len(deque)
		sum := 0
		for i := 0; i < curSize; i++ {
			sum += deque[i].Val
			if deque[i].Left != nil {
				deque = append(deque, deque[i].Left)
			}
			if deque[i].Right != nil {
				deque = append(deque, deque[i].Right)
			}
		}
		if curSize == len(deque) {
			return sum
		} else {
			deque = deque[curSize:]
		}
	}
	return -1
}
