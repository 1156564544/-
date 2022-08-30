package Question_by_day

func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parents[i]
		children[p] = append(children[p], i)
	}
	ans := 1
	max := 0
	var dfs func(int) int
	dfs = func(node int) int {
		res, size := 1, n-1
		for _, ch := range children[node] {
			sz := dfs(ch)
			size -= sz
			res *= sz
		}
		if node != 0 {
			res *= size
		}
		if res == max {
			ans += 1
		} else if res > max {
			max = res
			ans = 1
		}
		return n - size
	}
	dfs(0)
	return ans
}
