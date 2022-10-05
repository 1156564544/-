package Question_by_day

func getKthMagicNumber(k int) int {
	dp := make([]int, k+1)
	dp[1] = 1
	p1, p2, p3 := 1, 1, 1
	for i := 2; i <= k; i++ {
		next := min(dp[p1]*3, dp[p2]*5, dp[p3]*7)
		if next == dp[p1]*3 {
			p1 += 1
		}
		if next == dp[p2]*5 {
			p2 += 1
		}
		if next == dp[p3]*7 {
			p3 += 1
		}
		dp[i] = next
	}
	return dp[k]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else if b < c {
		return b
	} else {
		return c
	}
}
