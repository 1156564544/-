package Question_by_day

func parseBoolExpr(expression string) bool {
	stack := make([]byte, 0)
	for _, ch := range expression {
		if ch == ',' {
			continue
		}
		if ch == ')' {
			nT, nF := 0, 0
			j := len(stack) - 1
			for ; stack[j] != '('; j-- {
				if stack[j] == 't' {
					nT += 1
				}
				if stack[j] == 'f' {
					nF += 1
				}
			}
			if stack[j-1] == '!' {
				stack = stack[:j-1]
				push(&stack, nF == 1)
			} else if stack[j-1] == '&' {
				stack = stack[:j-1]
				push(&stack, nF == 0)
			} else if stack[j-1] == '|' {
				stack = stack[:j-1]
				push(&stack, nT != 0)
			}
		} else {
			stack = append(stack, byte(ch))
		}
	}

	if stack[0] == 't' {
		return true
	}
	return false
}

func push(stack *[]byte, flag bool) {
	if flag {
		*stack = append(*stack, 't')
	} else {
		*stack = append(*stack, 'f')
	}
}
