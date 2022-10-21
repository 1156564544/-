package Question_by_day

import "sort"

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	idx := make([]int, n)
	for i := 1; i < n; i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return endTime[idx[i]] < endTime[idx[j]]
	})
	dp := make([]int, n)
	dp[0] = profit[idx[0]]
	for i := 1; i < n; i++ {
		dp[i] = profit[idx[i]]
		for j := i - 1; j >= 0; j-- {
			if endTime[idx[j]] <= startTime[idx[i]] {
				dp[i] += dp[j]
				break
			}
		}
		dp[i] = max(dp[i], dp[i-1])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
