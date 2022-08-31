package Question_by_day

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if val > root.Val {
		return &TreeNode{val, root, nil}
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}
