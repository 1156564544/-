package Question_by_day

func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}
	if k == mi(n-1) {
		return '1'
	}
	if k < mi(n-1) {
		return findKthBit(n-1, k)
	} else {
		return invert(findKthBit(n-1, 2*mi(n-1)-k))
	}
}

func mi(a int) int {
	res := 1
	for i := 0; i < a; i++ {
		res = res * 2
	}
	return res
}

func invert(a byte) byte {
	if a == '1' {
		return '0'
	}
	return '1'
}
