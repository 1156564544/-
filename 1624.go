package Question_by_day

func maxLengthBetweenEqualCharacters(s string) int {
	m := make(map[byte]int)
	res := -1
	ss := []byte(s)
	for i, ch := range ss {
		if idx, ok := m[ch]; ok {
			res = max(res, i-idx-1)
		} else {
			m[ch] = i
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
