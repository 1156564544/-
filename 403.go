package Question_by_day

type pair struct {
	start int
	tiao  int
}

func canCross(stones []int) bool {
	if stones[1] != 1 {
		return false
	}
	n := len(stones)
	m := make(map[pair]bool)
	var dfs func(start, tiao int) bool
	dfs = func(start, tiao int) bool {
		if start == n-1 {
			return true
		}
		if start >= n {
			return false
		}
		if _, ok := m[pair{start, tiao}]; ok {
			return false
		}
		m[pair{start, tiao}] = false
		res := false
		for i := start + 1; i < n; i++ {
			if stones[i] == stones[start]+tiao-1 || stones[i] == stones[start]+tiao || stones[i] == stones[start]+tiao+1 {
				res = res || dfs(i, stones[i]-stones[start])
				if res == true {
					return true
				}
			} else if stones[i] > stones[start]+tiao+1 {
				break
			}
		}
		return false
	}
	return dfs(1, 1)
}
