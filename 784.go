package Question_by_day

func letterCasePermutation(s string) []string {
	res := make([]string, 1)
	res[0] = s
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' {
			l := len(res)
			for j := 0; j < l; j++ {
				res = append(res, swap(res[j], i))
			}

		}
	}
	return res
}

func swap(src string, idx int) (des string) {
	if src[idx] <= 'z' && src[idx] >= 'a' {
		des = src[:idx] + string(src[idx]+'A'-'a') + src[idx+1:]
	} else {
		des = src[:idx] + string(src[idx]+'a'-'A') + src[idx+1:]
	}
	return
}
