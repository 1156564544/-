package Question_by_day

func generateTheString(n int) string {
	ans := make([]byte, 0)
	if n%2 == 0 {
		for i := 0; i < n-1; i++ {
			ans = append(ans, 'a')
		}
		ans = append(ans, 'b')
	} else {
		for i := 0; i < n; i++ {
			ans = append(ans, 'a')
		}
	}
	return string(ans)
}
