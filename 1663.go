package Question_by_day

func getSmallestString(n int, k int) string {
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		cur := min(k-n+i+1, 26)
		res[n-i-1] = byte('a' + cur - 1)
		k = k - cur
		if cur == 1 {
			for j := i + 1; j < n; j++ {
				res[n-j-1] = 'a'
			}
			return string(res)
		}
	}
	return string(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
