package Question_by_day

func constructArray(n int, k int) []int {
	res := make([]int, n)
	for i := 0; i < n-k; i++ {
		res[i] = i + 1
	}

	for i := n - k; i < n; i += 2 {
		res[i] = n - (i-n+k)/2
	}
	for i := n - k + 1; i < n; i += 2 {
		res[i] = n - k + 1 + (i-n+k-1)/2
	}
	return res
}
