package Question_by_day

func numSimilarGroups(strs []string) int {
	n := len(strs)
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, 0)
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if help(strs[i], strs[j]) {
				m[i] = append(m[i], j)
				m[j] = append(m[j], i)
			}
		}
	}
	res := 0
	flag := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		flag[i] = true
		for _, j := range m[i] {
			if flag[j] == false {
				dfs(j)
			}
		}
	}
	for i := 0; i < n; i++ {
		if flag[i] == false {
			res += 1
			dfs(i)
		}
	}

	return res
}

func help(a, b string) bool {
	idx1, idx2 := -1, -1

	for i := range a {
		if a[i] != b[i] {
			if idx1 == -1 {
				idx1 = i
			} else if idx2 == -1 {
				idx2 = i
			} else {
				return false
			}
		}
	}

	return (idx1 == -1 && idx2 == -1) || (a[idx1] == b[idx2] && a[idx2] == b[idx1])
}
