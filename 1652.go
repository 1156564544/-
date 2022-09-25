package Question_by_day

func decrypt(code []int, k int) []int {
	n := len(code)
	if k == 0 {
		for i := 0; i < n; i++ {
			code[i] = 0
		}
	} else if k > 0 {
		for i := 0; i < k; i++ {
			code = append(code, code[i])
		}
		for i := 0; i < n; i++ {
			sum := 0
			for j := i + 1; j <= i+k; j++ {
				sum += code[j]
			}
			code[i] = sum
		}
		code = code[:n]
	} else {
		code = append(code, code[n+k:]...)
		for i := n - 1; i >= 0; i-- {
			sum := 0
			for j := i - 1; j >= i+k; j-- {
				sum += code[(j+n-k)%(n-k)]
			}
			code[i] = sum
		}
		code = code[:n]
	}
	return code
}
