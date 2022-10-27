package Question_by_day

import "fmt"

func longestValidParentheses(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 0
	res := 0
	for i := 1; i < n; i++ {
		if s[i] == '(' {
			dp[i] = 0
			continue
		}
		if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			dp[i] = dp[i-1] + 2
			if i-dp[i-1]-2 >= 0 {
				dp[i] += dp[i-dp[i-1]-2]
			}
			res = max(res, dp[i])
		}
	}
	fmt.Println(dp)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
