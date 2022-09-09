package Question_by_day

func numMatchingSubseq(s string, words []string) int {
	res := 0
	ss := []byte(s)
	for _, w := range words {
		if help(ss, []byte(w)) {
			res += 1
		}
	}
	return res
}

func help(a, b []byte) bool {
	if len(b) == 0 {
		return true
	}
	if len(a) < len(b) {
		return false
	}
	for i, c := range a {
		if c == b[0] {
			return help(a[i+1:], b[1:])
		}
	}
	return false
}
