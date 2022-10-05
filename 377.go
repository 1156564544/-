package Question_by_day

import "sort"

func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, j := range nums {
			if j <= i {
				dp[i] += dp[i-j]
			}
		}
	}
	return dp[target]
}
