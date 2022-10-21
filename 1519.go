package Question_by_day

func countSubTrees(n int, edges [][]int, labels string) []int {
	houxu := make([][]int, n)
	for i := 0; i < n; i++ {
		houxu[i] = make([]int, 0)
	}
	for _, e := range edges {
		houxu[e[0]] = append(houxu[e[0]], e[1])
		houxu[e[1]] = append(houxu[e[1]], e[0])
	}
	res := make([]int, n)
	var dfs func(idx int) []int
	dfs = func(idx int) []int {
		res[idx] = 1
		m := make([]int, 26)
		m[labels[idx]-'a'] = 1
		for _, i := range houxu[idx] {
			if res[i] != 0 {
				continue
			}
			mm := dfs(i)
			for k, v := range mm {
				m[k] += v
			}
		}
		res[idx] = m[labels[idx]-'a']
		return m
	}
	dfs(0)

	return res
}
