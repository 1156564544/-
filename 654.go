package Question_by_day

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}
	idx := 0
	max := nums[0]
	for i, v := range nums {
		if max < v {
			max = v
			idx = i
		}
	}
	root := &TreeNode{Val: max}
	root.Left = constructMaximumBinaryTree(nums[:idx])
	root.Right = constructMaximumBinaryTree(nums[idx+1:])
	return root
}
