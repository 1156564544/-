package Question_by_day

func minAddToMakeValid(s string) int {
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left += 1
		} else {
			if left > 0 {
				left -= 1
			} else {
				right += 1
			}
		}
	}
	return left + right
}
