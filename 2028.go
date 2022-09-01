package Question_by_day

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	total := mean * (m + n)
	for _, r := range rolls {
		total -= r
	}
	if total < n || total > 6*n {
		return []int{}
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		x := total / (n - i)
		if total == x*(n-i) {
			res[i] = x
			total = total - x
		} else {
			res[i] = x + 1
			total = total - x - 1
		}
	}
	return res
}
