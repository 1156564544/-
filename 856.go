package Question_by_day

func scoreOfParentheses(s string) int {
	stack := []int{}
	ss := []byte(s)
	for _, ch := range ss {
		if ch == '(' {
			stack = append(stack, 0)
		} else {
			i := len(stack) - 1
			sum := 0
			for ; i >= 0; i-- {
				if stack[i] == 0 {
					break
				} else {
					sum += stack[i]
				}
			}
			if sum != 0 {
				sum = sum * 2
			} else {
				sum = 1
			}
			stack = stack[:i]
			stack = append(stack, sum)
		}
	}
	sum := 0
	for _, s := range stack {
		sum += s
	}
	return sum
}
