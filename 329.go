package Question_by_day

func longestIncreasingPath(matrix [][]int) int {
	res := 1
	m := len(matrix)
	n := len(matrix[0])
	c := make([][]int, m)
	for i := 0; i < m; i++ {
		c[i] = make([]int, n)
	}
	var check func(i, j int) bool
	check = func(i, j int) bool {
		return !(i < 0 || i >= m || j < 0 || j >= n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if c[i][j] > 0 {
			return c[i][j]
		}
		var res int = 0
		if check(i-1, j) && matrix[i-1][j] > matrix[i][j] {
			res = max(res, dfs(i-1, j))
		}
		if check(i+1, j) && matrix[i+1][j] > matrix[i][j] {
			res = max(res, dfs(i+1, j))
		}
		if check(i, j-1) && matrix[i][j-1] > matrix[i][j] {
			res = max(res, dfs(i, j-1))
		}
		if check(i, j+1) && matrix[i][j+1] > matrix[i][j] {
			res = max(res, dfs(i, j+1))
		}
		c[i][j] = 1 + res
		return 1 + res
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(res, dfs(i, j))
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
