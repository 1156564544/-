package Question_by_day

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	if k > mi(n-2) {
		return 1 ^ kthGrammar(n-1, k-mi(n-2))
	} else {
		return kthGrammar(n-1, k)
	}
}

func mi(n int) int {
	res := 1
	for n > 0 {
		res = res * 2
		n -= 1
	}
	return res
}
