package Question_by_day

type pair struct {
	x int
	y int
}

func shortestBridge(grid [][]int) (ans int) {
	n := len(grid)
	des := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				island := []pair{pair{i, j}}
				q := []pair{pair{i, j}}
				grid[i][j] = -1
				for len(q) > 0 {
					p := q[0]
					q = q[1:]
					for _, d := range des {
						if p.x+d[0] >= 0 && p.x+d[0] < n && p.y+d[1] >= 0 && p.y+d[1] < n && grid[p.x+d[0]][p.y+d[1]] == 1 {
							island = append(island, pair{p.x + d[0], p.y + d[1]})
							q = append(q, pair{p.x + d[0], p.y + d[1]})
							grid[p.x+d[0]][p.y+d[1]] = -1
						}
					}
				}
				for true {
					l := len(island)
					for i := 0; i < l; i++ {
						p := island[i]
						for _, d := range des {
							if p.x+d[0] >= 0 && p.x+d[0] < n && p.y+d[1] >= 0 && p.y+d[1] < n {
								if grid[p.x+d[0]][p.y+d[1]] == 1 {
									return
								}
								if grid[p.x+d[0]][p.y+d[1]] == 0 {
									grid[p.x+d[0]][p.y+d[1]] = -1
									island = append(island, pair{p.x + d[0], p.y + d[1]})
								}
							}
						}
					}
					ans += 1
				}
			}
		}
	}
	return
}
