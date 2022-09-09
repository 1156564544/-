package Question_by_day

func reorderSpaces(text string) string {
	ss := []byte(text)
	words := []string{}
	idx := -1
	sum := 0
	for i := 0; i < len(ss); i++ {
		if ss[i] == ' ' {
			sum += 1
			if i > 0 && ss[i-1] != ' ' {
				words = append(words, string(ss[idx:i]))
			}
		} else {
			if i == 0 || ss[i-1] == ' ' {
				idx = i
			}
		}
	}
	if ss[len(ss)-1] != ' ' {
		words = append(words, string(ss[idx:]))
	}

	res := ""
	if len(words) == 1 {
		res = words[0]
		for i := 0; i < sum; i++ {
			res = res + " "
		}
	} else {
		num := sum / (len(words) - 1)
		tap := ""
		for i := 0; i < num; i++ {
			tap = tap + " "
		}
		for i := 0; i < len(words)-1; i++ {
			res = res + words[i] + tap
		}
		res = res + words[len(words)-1]
		for i := num * (len(words) - 1); i < sum; i++ {
			res = res + " "
		}
	}
	return res
}
