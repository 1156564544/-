package Question_by_day

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	res := n + 1
	q := make([]int, 0)
	q = append(q, 0)
	idx := make([]int, 0)
	idx = append(idx, 0)
	for i := 1; i <= n; i++ {
		var j int
		for j = len(q) - 1; j >= 0; j-- {
			if q[j] < sum[i] {
				break
			}
		}

		l := 0
		for l = 0; l <= j; l++ {
			if sum[i]-q[l] < k {
				break
			}
		}
		if l > 0 {
			res = min(res, i-idx[l-1])
		}
		q = q[l : j+1]
		idx = idx[l : j+1]
		q = append(q, sum[i])
		idx = append(idx, i)
	}

	if res == n+1 {
		return -1
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
