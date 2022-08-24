package Question_by_day

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return help2(root) - 1
}

func help2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	if root.Left != nil && root.Left.Val == root.Val {
		res += help(root.Left)
	}
	if root.Right != nil && root.Right.Val == root.Val {
		res += help(root.Right)
	}
	l := help2(root.Left)
	r := help2(root.Right)
	res = maxF(res, maxF(l, r))

	return res
}

func help(root *TreeNode) int {
	var a, b int
	if root.Left != nil && root.Left.Val == root.Val {
		a = help(root.Left)
	}
	if root.Right != nil && root.Right.Val == root.Val {
		b = help(root.Right)
	}
	return 1 + maxF(a, b)
}

func maxF(a, b int) int {
	if a > b {
		return a
	}
	return b
}
