package Question_by_day

func minWindow(s string, t string) string {
	output := ""
	n := len(s)
	leng := n + 1
	l, r := 0, 0
	m := make(map[byte]int)
	for _, b := range []byte(t) {
		m[b] += 1
	}
	m[s[0]] -= 1
	for r < n {
		if help(m) {
			if r-l+1 < leng {
				output = s[l : r+1]
				leng = r - l + 1
			}
			m[s[l]] += 1
			l += 1
		} else {
			r += 1
			if r == n {
				break
			}
			m[s[r]] -= 1
		}
	}
	return output
}

func help(m map[byte]int) bool {
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}
