package Question_by_day

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	m := make(map[byte]int)
	for _, ch := range []byte(s1) {
		m[ch] += 1
	}
	ss := []byte(s2)
	n := len(s1)
	for i := 0; i < n; i++ {
		m[ss[i]] -= 1
		if m[ss[i]] == 0 {
			delete(m, ss[i])
		}
	}
	for i := n; i < len(ss); i++ {
		if len(m) == 0 {
			return true
		}
		m[ss[i]] -= 1
		if m[ss[i]] == 0 {
			delete(m, ss[i])
		}
		m[ss[i-n]] += 1
		if m[ss[i-n]] == 0 {
			delete(m, ss[i-n])
		}
	}
	return len(m) == 0
}
