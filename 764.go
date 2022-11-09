package Question_by_day

func orderOfLargestPlusSign(n int, mines [][]int) int {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			grid[i][j] = 1
		}
	}
	for _, node := range mines {
		grid[node[0]][node[1]] = 0
	}
	c1 := make([][]int, n)
	c2 := make([][]int, n)
	for i := 0; i < n; i++ {
		c1[i] = make([]int, n)
		c2[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				c1[i][j] = 0
			} else {
				if grid[i-1][j] == 1 {
					c1[i][j] = c1[i-1][j] + 1
				}
			}
			if j == 0 {
				c2[i][j] = 0
			} else {
				if grid[i][j-1] == 1 {
					c2[i][j] = c2[i][j-1] + 1
				}
			}
		}
	}
	d1 := make([][]int, n)
	d2 := make([][]int, n)
	for i := 0; i < n; i++ {
		d1[i] = make([]int, n)
		d2[i] = make([]int, n)
	}
	res := 0
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == n-1 {
				d1[i][j] = 0
			} else {
				if grid[i+1][j] == 1 {
					d1[i][j] = d1[i+1][j] + 1
				}
			}
			if j == n-1 {
				d2[i][j] = 0
			} else {
				if grid[i][j+1] == 1 {
					d2[i][j] = d2[i][j+1] + 1
				}
			}
			if grid[i][j] == 1 {
				x := min(c1[i][j], c2[i][j], d1[i][j], d2[i][j])
				res = max(res, x+1)
			}
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
func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func min(a, b, c, d int) int {
	return min2(min2(a, b), min2(c, d))
}
