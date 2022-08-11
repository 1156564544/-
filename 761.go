package Question_by_day

import "sort"

func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}
	left, cnt := 0, 0
	ans := make([]string, 0)
	for i, ch := range s {
		if ch == '1' {
			cnt += 1
		} else {
			if cnt--; cnt == 0 {
				ans = append(ans, "1"+makeLargestSpecial(s[left+1:i])+"0")
				left = i + 1
			}
		}
	}
	sort.Strings(ans)
	res := ""
	for _, ss := range ans {
		res = ss + res
	}
	return res
}
