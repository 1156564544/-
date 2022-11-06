package Question_by_day

func numRollsToTarget(n int, k int, target int) int {
	if k > target {
		k = target
	}
	dp := make([]int64, target+1)
	temp := make([]int64, target+1)
	for i := 1; i <= k; i++ {
		dp[i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := i - 1; j <= target-(n-i+1); j++ {
			for l := 1; l <= k && j+l <= target-(n-i); l++ {
				// temp[j+l]=(temp[j+1]+dp[j])%1000000007
				temp[j+l] += dp[j] % 1000000007
			}
		}
		copy(dp, temp)
		for id := 1; id <= target; id++ {
			temp[id] = 0
		}
	}
	return int(dp[target] % 1000000007)
}
