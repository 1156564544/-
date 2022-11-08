package Question_by_day

func winnerSquareGame(n int) bool {
	count := make([]int, 1)
	count[0] = 1
	for i := 2; i*i <= n; i++ {
		count = append(count, i*i)
	}
	num := make([]bool, n+1)
	num[1] = true
	for i := 2; i <= n; i++ {
		if num[i] {
			continue
		}
		for _, v := range count {
			if v > i {
				break
			}
			if num[i-v] == false {
				num[i] = true
				break
			}
		}
		if !num[i] {
			for _, v := range count {
				if i+v > n {
					break
				}
				num[i+v] = true
			}
		}
	}
	return num[n]
}
