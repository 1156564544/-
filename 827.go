package Question_by_day

type pair struct {
	x int
	y int
}

func largestIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	cur := -1
	count := make([]int, 0)
	mm := make(map[pair]int)
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] != 1 {
			return
		}
		grid[i][j] = 2
		count[cur] += 1
		mm[pair{i, j}] = cur
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				cur += 1
				count = append(count, 0)
				dfs(i, j)
				if count[cur] > max {
					max = count[cur]
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans := 1
			dao := make(map[int]bool)
			if grid[i][j] == 0 {
				if i > 0 && grid[i-1][j] == 2 {
					id := mm[pair{i - 1, j}]
					ans += count[id]
					dao[id] = true
				}
				if i < m-1 && grid[i+1][j] == 2 {
					id := mm[pair{i + 1, j}]
					if dao[id] == false {
						dao[id] = true
						ans += count[id]
					}
				}
				if j > 0 && grid[i][j-1] == 2 {
					id := mm[pair{i, j - 1}]
					if dao[id] == false {
						dao[id] = true
						ans += count[id]
					}
				}
				if j < n-1 && grid[i][j+1] == 2 {
					id := mm[pair{i, j + 1}]
					if dao[id] == false {
						dao[id] = true
						ans += count[id]
					}
				}
			}
			if max < ans {
				max = ans
			}
		}
	}
	return max
}
