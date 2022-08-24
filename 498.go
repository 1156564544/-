package Question_by_day

func findDiagonalOrder(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])
	res := make([]int, m*n)
	i, j := 0, 0
	flag := true
	for cnt := 0; cnt < m*n; cnt++ {
		res[cnt] = mat[i][j]
		if flag {
			i -= 1
			j += 1
			if j == n {
				i += 2
				j -= 1
				flag = false
				continue
			}
			if i < 0 {
				i = 0
				flag = false
			}
		} else {
			i += 1
			j -= 1
			if i == m {
				i -= 1
				j += 2
				flag = true
				continue
			}
			if j < 0 {
				j = 0
				flag = true
			}
		}
	}
	return res
}
