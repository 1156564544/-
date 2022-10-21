package Question_by_day

func construct(board [][]int) [6]int {
	res := [6]int{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			res[i*3+j] = board[i][j]
		}
	}
	return res
}

type status struct {
	board [6]int
	count int
}

func slidingPuzzle(board [][]int) int {
	m := make(map[[6]int]bool)
	q := make([]status, 1)
	q[0] = status{construct(board), 0}
	for len(q) > 0 {
		cur_board := q[0].board
		cur_count := q[0].count
		if cur_board == [6]int{1, 2, 3, 4, 5, 0} {
			return cur_count
		}
		q = q[1:]
		if m[cur_board] {
			continue
		}
		m[cur_board] = true
		var i, j int
	Loop:
		for i = 0; i < 2; i++ {
			for j = 0; j < 3; j++ {
				if cur_board[i*3+j] == 0 {
					break Loop
				}
			}
		}
		if i > 0 {
			cur_board[i*3+j], cur_board[(i-1)*3+j] = cur_board[(i-1)*3+j], cur_board[i*3+j]
			q = append(q, status{cur_board, cur_count + 1})
			cur_board[i*3+j], cur_board[(i-1)*3+j] = cur_board[(i-1)*3+j], cur_board[i*3+j]
		}
		if i < 1 {
			cur_board[i*3+j], cur_board[(i+1)*3+j] = cur_board[(i+1)*3+j], cur_board[i*3+j]
			q = append(q, status{cur_board, cur_count + 1})
			cur_board[i*3+j], cur_board[(i+1)*3+j] = cur_board[(i+1)*3+j], cur_board[i*3+j]
		}
		if j > 0 {
			cur_board[i*3+j], cur_board[i*3+j-1] = cur_board[i*3+j-1], cur_board[i*3+j]
			q = append(q, status{cur_board, cur_count + 1})
			cur_board[i*3+j], cur_board[i*3+j-1] = cur_board[i*3+j-1], cur_board[i*3+j]
		}
		if j < 2 {
			cur_board[i*3+j], cur_board[i*3+j+1] = cur_board[i*3+j+1], cur_board[i*3+j]
			q = append(q, status{cur_board, cur_count + 1})
			cur_board[i*3+j], cur_board[i*3+j+1] = cur_board[i*3+j+1], cur_board[i*3+j]
		}
	}

	return -1
}
