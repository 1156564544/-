package Question_by_day

func maxProduct(root *TreeNode) int {
	m := make(map[*TreeNode]int)
	var dfs1 func(*TreeNode)
	dfs1 = func(root *TreeNode) {
		m[root] = root.Val
		if root.Left != nil {
			dfs1(root.Left)
			m[root] += m[root.Left]
		}
		if root.Right != nil {
			dfs1(root.Right)
			m[root] += m[root.Right]
		}
	}
	dfs1(root)
	sum := m[root]
	var dfs2 func(*TreeNode) int
	dfs2 = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		best := m[root]
		l := dfs2(root.Left)
		r := dfs2(root.Right)
		if abs(2*l-sum) < abs(2*best-sum) {
			best = l
		}
		if abs(2*r-sum) < abs(2*best-sum) {
			best = r
		}
		return best
	}
	best := dfs2(root)
	var res int64
	res = int64(best) * int64(sum-best)
	return int(res % 1000000007)
}
func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
