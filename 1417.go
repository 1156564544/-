package Question_by_day

func reformat(s string) string {
	chars := []byte{}
	numbers := []byte{}
	ss := []byte(s)
	for _, ch := range ss {
		if ch >= '0' && ch <= '9' {
			numbers = append(numbers, ch)
		} else {
			chars = append(chars, ch)
		}
	}
	res := []byte{}
	if len(numbers) == len(chars) || len(numbers) == len(chars)+1 || len(numbers) == len(chars)-1 {
		min := len(numbers)
		if min > len(chars) {
			min = len(chars)
		}
		for i := 0; i < min; i++ {
			res = append(res, chars[i], numbers[i])
		}
		if len(numbers) > min {
			return string(append([]byte{numbers[min]}, res...))
		}
		if len(chars) > min {
			return string(append(res, chars[min]))
		}
		return string(res)
	}
	return ""
}
