package Question_by_day

func numWays(n int) int {
	if n < 2 {
		return 1
	}
	x := help(n - 1)
	return int((x[0][0]%(1e9+7) + x[0][1]%(1e9+7)) % (1e9 + 7))
}

func help(n int) [][]int64 {
	if n == 0 {
		return [][]int64{{1, 0}, {0, 1}}
	}
	if n == 1 {
		return [][]int64{{1, 1}, {1, 0}}
	}
	x1 := help(n / 2)
	x2 := help(n - n/2)
	res := make([][]int64, 2)
	for i := 0; i < 2; i++ {
		res[i] = make([]int64, 2)
	}
	res[0][0] = (x1[0][0]*x2[0][0] + x1[0][1]*x2[1][0]) % (1e9 + 7)
	res[0][1] = (x1[0][0]*x2[0][1] + x1[0][1]*x2[1][1]) % (1e9 + 7)
	res[1][0] = (x1[1][0]*x2[0][0] + x1[1][1]*x2[1][0]) % (1e9 + 7)
	res[1][1] = (x1[1][0]*x2[0][1] + x1[1][1]*x2[1][1]) % (1e9 + 7)
	return res
}
