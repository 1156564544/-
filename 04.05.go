package Question_by_day

import "math"

func isValidBST(root *TreeNode) bool {
	return help(root, math.MinInt64, math.MaxInt64)
}
func help(root *TreeNode, low, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val >= upper || root.Val <= low {
		return false
	}
	return help(root.Left, low, root.Val) && help(root.Right, root.Val, upper)
}
