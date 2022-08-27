package Question_by_day

func widthOfBinaryTree(root *TreeNode) int {
	q1 := make([]*TreeNode, 0)
	res := 1
	q2 := make([]int, 0)
	if root.Left != nil {
		q1 = append(q1, root.Left)
		q2 = append(q2, 0)
	}
	if root.Right != nil {
		q1 = append(q1, root.Right)
		q2 = append(q2, 1)
	}
	for len(q1) > 0 {
		n := len(q1)
		left := q2[0]
		right := q2[n-1]
		if right-left+1 > res {
			res = right - left + 1
		}
		for i := 0; i < n; i++ {
			if q1[i].Left != nil {
				q1 = append(q1, q1[i].Left)
				q2 = append(q2, q2[i]*2)
			}
			if q1[i].Right != nil {
				q1 = append(q1, q1[i].Right)
				q2 = append(q2, q2[i]*2+1)
			}
		}
		q1 = q1[n:]
		q2 = q2[n:]
	}
	return res
}
