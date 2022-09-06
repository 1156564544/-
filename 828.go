package Question_by_day

func uniqueLetterString(s string) int {
	res := 0
	ss := []byte(s)
	n := len(ss)
	m := make(map[byte][]int)
	for i, c := range ss {
		m[c] = append(m[c], i)
	}
	for _, v := range m {
		l := -1
		for i := 0; i < len(v)-1; i++ {
			res += (v[i] - l) * (v[i+1] - v[i])
			l = v[i]
		}
		res += (v[len(v)-1] - l) * (n - v[len(v)-1])
	}

	return res
}
