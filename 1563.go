package Question_by_day

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	dp := make([][]int, n)
	sum := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		sum[i] = make([]int, n)
		sum[i][i] = stoneValue[i]
	}
	for len := 2; len <= n; len++ {
		for i := 0; i+len <= n; i++ {
			sum[i][i+len-1] = sum[i][i+len-2] + stoneValue[i+len-1]
		}
	}
	for len := 2; len <= n; len++ {
		for l := 0; l+len-1 < n; l++ {
			r := l + len - 1
			for i := l; i < r; i++ {
				ls := sum[l][i]
				rs := sum[i+1][r]
				if ls < rs {
					dp[l][r] = max(dp[l][r], dp[l][i]+sum[l][i])
				} else if ls > rs {
					dp[l][r] = max(dp[l][r], dp[i+1][r]+sum[i+1][r])
				} else {
					dp[l][r] = max(dp[l][r], max(dp[l][i], dp[i+1][r])+sum[i+1][r])
				}
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
