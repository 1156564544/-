package Question_by_day

func shiftingLetters(s string, shifts []int) string {
	n := len(shifts)
	res := []byte(s)
	for i := n - 2; i >= 0; i-- {
		shifts[i] += shifts[i+1]
	}
	for i := 0; i < n; i++ {
		res[i] = 'a' + byte((int(res[i]-'a')+shifts[i])%26)
	}
	return string(res)
}
