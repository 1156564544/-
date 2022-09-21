package Question_by_day

func isMatch(s string, p string) bool {
	ss := []byte(s)
	pp := []byte(p)
	n := len(ss)
	m := len(pp)
	match := func(i, j int) bool {
		if pp[i] == '.' || pp[i] == ss[j] {
			return true
		} else {
			return false
		}
	}
	c := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		c[i] = make([]bool, n+1)
	}
	c[0][0] = true
	for i := 1; i <= n; i++ {
		c[0][i] = false
	}
	for i := 1; i <= m; i++ {
		if pp[i-1] == '*' {
			if i == 2 || c[i-2][0] {
				c[i][0] = true
			}
		}
		for j := 1; j <= n; j++ {
			if pp[i-1] == '*' {
				c[i][j] = c[i-2][j]
				for k := j - 1; k >= 0; k-- {
					if match(i-2, k) {
						c[i][j] = c[i][j] || c[i-2][k]
						if c[i][j] {
							break
						}
					} else {
						break
					}
				}
			} else if match(i-1, j-1) {
				c[i][j] = c[i-1][j-1]
			}
		}
	}
	return c[m][n]
}
