package Question_by_day

func numSpecial(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	rows := make([]int, m)
	cols := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				rows[i] += 1
				cols[j] += 1
			}
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		if rows[i] != 1 {
			continue
		}
		for j := 0; j < n; j++ {
			if cols[j] == 1 && mat[i][j] == 1 {
				res += 1
				break
			}
		}
	}
	return res
}
