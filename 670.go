package Question_by_day

func maximumSwap(num int) int {
	idx := 0
	x := num
	for x > 0 {
		idx += 1
		x = x / 10
	}
	res := num
	for i := idx; i > 1; i-- {
		a := help(num, i)
		max := help(num, i-1)
		k := i - 1
		for j := i - 1; j > 0; j-- {
			b := help(num, j)
			if b >= max {
				max = b
				k = j
			}
		}
		if a < max {
			res -= a * mi(i)
			res += max * mi(i)
			res -= max * mi(k)
			res += a * mi(k)
			return res
		}
	}
	return res
}

func help(num, idx int) int {
	num = num / mi(idx)
	return num % 10
}

func mi(idx int) int {
	res := 1
	for idx > 1 {
		res = res * 10
		idx -= 1
	}
	return res
}
