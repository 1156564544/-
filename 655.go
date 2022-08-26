package Question_by_day

import "strconv"

func printTree(root *TreeNode) [][]string {
	h := height(root) - 1
	res := make([][]string, h+1)
	n := mi(2, h+1) - 1
	for i := 0; i <= h; i++ {
		res[i] = make([]string, n)
	}

	help(root, 0, (n-1)/2, h, res)
	return res
}

func help(root *TreeNode, i, j, h int, res [][]string) {
	if root == nil {
		return
	}
	res[i][j] = strconv.Itoa(root.Val)
	help(root.Left, i+1, j-mi(2, h-i-1), h, res)
	help(root.Right, i+1, j+mi(2, h-i-1), h, res)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(height(root.Left), height(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mi(a, b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res *= a
	}
	return res
}
