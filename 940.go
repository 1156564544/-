package Question_by_day

func distinctSubseqII(s string) int {
	var chashu int64 = 1000000007
	n := len(s)
	c1, c2 := make([]int64, n), make([]int64, n)
	c1[0] = 0
	c2[0] = 1
	for i := 1; i < n; i++ {
		c1[i] = c1[i-1] + c2[i-1]
		c2[i] = c1[i] + 1
		for j := i - 1; j >= 0; j-- {
			if s[j] == s[i] {
				c2[i] = (c2[i] - c2[j] + chashu) % chashu
			}
		}
	}
	return int((c1[n-1] + c2[n-1]) % 1000000007)
}
