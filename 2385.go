package Question_by_day

func amountOfTime(root *TreeNode, start int) int {
	var dfs func(r *TreeNode) (int, int, int)
	dfs = func(r *TreeNode) (int, int, int) {
		if r == nil {
			return 0, -1, 0
		}
		lh, lw, lt := dfs(r.Left)
		rh, rw, rt := dfs(r.Right)
		if r.Val == start {
			return max(lh, rh) + 1, 0, max(lh, rh)
		}
		if lw != -1 || rw != -1 {
			if lw != -1 {
				return max(lh, rh) + 1, lw + 1, max(lw+1+rh, lt)
			} else {
				return max(lh, rh) + 1, rw + 1, max(rw+1+lh, rt)
			}
		} else {
			return max(lh, rh) + 1, -1, 0
		}
	}
	_, _, t := dfs(root)
	return t

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
