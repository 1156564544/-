package Question_by_day

func preimageSizeFZF(k int) int {
	sum := 0
	sum2 := 0
	for i := 25; sum < k; i += 25 {
		j := i / 25
		x := 6
		for j > 0 && j%5 == 0 {
			j = j / 5
			x += 1
		}
		sum2 = sum + 4
		sum += x
	}
	if k == sum || k <= sum2 {
		return 5
	} else {
		return 0
	}
}
