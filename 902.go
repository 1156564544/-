package Question_by_day

import "strconv"

func atMostNGivenDigitSet(digits []string, n int) int {
	num := len(digits)
	digit := make([]int, num)
	for i, v := range digits {
		digit[i], _ = strconv.Atoi(v)
	}
	res := 0
	w := ws(n)
	for i := 1; i < w; i++ {
		res += mi(num, i)
	}
	var help func(n, w int) int
	help = func(n, w int) int {
		if w == 0 {
			return 1
		}
		ans := 0
		v := n / mi(10, w-1)
		for i := 0; i < len(digit); i++ {
			if digit[i] < v {
				ans += mi(num, w-1)
			} else if digit[i] == v {
				ans += help(n-v*mi(10, w-1), w-1)
			} else {
				break
			}
		}
		return ans
	}

	return res + help(n, w)
}

func mi(a, b int) int {
	res := 1
	for i := 1; i <= b; i++ {
		res = res * a
	}
	return res
}

func ws(n int) int {
	res := 0
	for n > 0 {
		res += 1
		n = n / 10
	}
	return res
}
