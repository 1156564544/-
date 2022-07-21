package Question_by_day

func shiftGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := (i*n + j + k) % (m * n)
			x := idx / n
			y := idx - x*n
			res[x][y] = grid[i][j]
		}
	}

	return res
}
