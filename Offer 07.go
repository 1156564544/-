package Question_by_day

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{}
	root.Val = preorder[0]
	idx := 0
	for i, v := range inorder {
		if v == root.Val {
			idx = i
			break
		}
	}
	root.Left = buildTree(preorder[1:1+idx], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}
