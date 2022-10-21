package Question_by_day

func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n)
	for _, d := range dislikes {
		g[d[0]-1] = append(g[d[0]-1], d[1]-1)
		g[d[1]-1] = append(g[d[1]-1], d[0]-1)
	}
	visit := make([]int, n)
	var dfs func(cur, c int) bool
	dfs = func(cur, c int) bool {
		visit[cur] = c
		for _, next := range g[cur] {
			if visit[next] == c || visit[next] == 0 && !dfs(next, 3^c) {
				return false
			}
		}
		return true
	}
	for i, c := range visit {
		if c == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}
