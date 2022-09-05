package Question_by_day

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	r := make([]int, n)
	r[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		r[i] = r[i+1] + nums[i]
	}
	max := r[n-1]
	for i := n - 1; i >= 0; i-- {
		if r[i] < max {
			r[i] = max
		}
		if r[i] > max {
			max = r[i]
		}
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < n; i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	sum := 0
	for i := 0; i < n-2; i++ {
		sum += nums[i]
		res = maxf(res, sum+r[i+2])
	}
	return res
}

func maxf(a, b int) int {
	if a < b {
		return b
	}
	return a
}
