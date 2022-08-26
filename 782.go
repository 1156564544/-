package Question_by_day

func movesToChessboard(board [][]int) int {
	n := len(board)
	firstRow := make([]int, n)
	firstRow = board[0]
	reverseRow := make([]int, n)
	for i, v := range firstRow {
		if v == 0 {
			reverseRow[i] = 1
		} else {
			reverseRow[i] = 0
		}
	}
	firstColumn := make([]int, n)
	reverseColumn := make([]int, n)
	for i := 0; i < n; i++ {
		if board[i][0] == 1 {
			firstColumn[i] = 1
			reverseColumn[i] = 0
		} else {
			firstColumn[i] = 0
			reverseColumn[i] = 1
		}
	}
	for i := 0; i < n; i++ {
		flag := true
		for j := 0; j < n; j++ {
			if board[i][j] != firstRow[j] {
				flag = false
				break
			}
		}
		if flag == true {
			continue
		}
		for j := 0; j < n; j++ {
			if board[i][j] != reverseRow[j] {
				return -1
			}
		}
	}
	for i := 0; i < n; i++ {
		flag := true
		for j := 0; j < n; j++ {
			if board[j][i] != firstColumn[j] {
				flag = false
				break
			}
		}
		if flag == true {
			continue
		}
		for j := 0; j < n; j++ {
			if board[j][i] != reverseColumn[j] {
				return -1
			}
		}
	}
	res := 0
	num_row1 := 0
	num_column1 := 0
	for _, v := range firstRow {
		if v == 1 {
			num_row1 += 1
		}
	}
	for _, v := range firstColumn {
		if v == 1 {
			num_column1 += 1
		}
	}
	jishu_row1 := 0
	jishu_column1 := 0
	for i := 0; i < n; i += 2 {
		if firstRow[i] == 1 {
			jishu_row1 += 1
		}
	}
	for i := 0; i < n; i += 2 {
		if firstColumn[i] == 1 {
			jishu_column1 += 1
		}
	}
	if n%2 == 0 {
		if n != 2*num_row1 || n != 2*num_column1 {
			return -1
		}
		res += min(jishu_row1, n/2-jishu_row1) + min(jishu_column1, n/2-jishu_column1)
		return res
	}
	if (n != 2*num_column1-1 && n != 2*num_column1+1) || (n != 2*num_row1-1 && n != 2*num_row1+1) {
		return -1
	}
	if n == 2*num_column1-1 {
		res += num_column1 - jishu_column1
	} else {
		res += jishu_column1
	}
	if n == 2*num_row1-1 {
		res += num_row1 - jishu_row1
	} else {
		res += jishu_row1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
