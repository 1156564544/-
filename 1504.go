package Question_by_day

type pair struct {
	height int
	column int
}

func numSubmat(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	dat := make([][]int, m)
	for i := 0; i < m; i++ {
		dat[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if mat[i][0] == 1 {
			dat[i][0] = 1
		}
		for j := 1; j < n; j++ {
			if mat[i][j] == 0 {
				continue
			}
			dat[i][j] = dat[i][j-1]
			if mat[i][j] == 1 {
				dat[i][j] += 1
			}
		}
	}
	res := 0
	for j := 0; j < n; j++ {
		stack := make([]pair, 1)
		stack[0] = pair{1, dat[0][j]}
		sum := dat[0][j]
		res += sum
		for i := 1; i < m; i++ {
			height := 1
			sum += dat[i][j]
			for len(stack) > 0 && stack[len(stack)-1].column > dat[i][j] {
				sum -= (stack[len(stack)-1].column - dat[i][j]) * stack[len(stack)-1].height
				height += stack[len(stack)-1].height
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, pair{height, dat[i][j]})
			res += sum
		}
	}

	return res
}
