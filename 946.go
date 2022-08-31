package Question_by_day

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	stack := make([]int, 0)
	i, j := 0, 0
	for i < n || j < n {
		if i < n {
			stack = append(stack, pushed[i])
			i += 1
		} else {
			return false
		}
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j += 1
		}
	}
	return true
}
