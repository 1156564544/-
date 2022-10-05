package Question_by_day

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, s := range cpdomains {
		i := strings.IndexByte(s, ' ')
		num, _ := strconv.Atoi(s[:i])
		s = s[i+1:]
		for {
			m[s] += num
			i := strings.IndexByte(s, '.')
			if i < 0 {
				break
			}
			s = s[i+1:]
		}
	}
	res := []string{}
	for k, v := range m {
		res = append(res, strconv.Itoa(v)+" "+k)
	}
	return res
}
