package Question_by_day

import (
	"strconv"
	"strings"
)

type stack struct {
	id    int
	opt   string
	time  int
	ideal int
}

func exclusiveTime(n int, logs []string) (ans []int) {
	res := make(map[int]int)
	s := []stack{}
	for _, l := range logs {
		log := strings.Split(l, ":")
		id, _ := strconv.Atoi(log[0])
		opt := log[1]
		time, _ := strconv.Atoi(log[2])
		if opt == "start" {
			s = append(s, stack{id, opt, time, 0})
		} else {
			n := len(s)
			start := s[n-1]
			s = s[:n-1]
			res[start.id] += time - start.time + 1 - start.ideal
			if len(s) > 0 {
				n = len(s)
				if s[n-1].opt == "start" {
					s[n-1].ideal += time - start.time + 1
				}
			}
		}

	}

	ans = make([]int, n)
	for k, v := range res {
		ans[k] = v
	}

	return
}
