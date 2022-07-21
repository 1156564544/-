package Question_by_day

func pruneTree(root *TreeNode) *TreeNode {
	flag := help(root)
	if flag == false {
		return nil
	}
	return root
}

func help(root *TreeNode) bool {
	if root == nil {
		return false
	}
	l := help(root.Left)
	r := help(root.Right)
	if l == false {
		root.Left = nil
	}
	if r == false {
		root.Right = nil
	}
	if root.Val == 0 {
		return l || r
	}
	return true
}
