package Question_by_day

import "fmt"

func minOperations(logs []string) int {
	cur := make([]string, 0)
	for _, log := range logs {
		cmd := strings.Split(log, "/")
		cmd = cmd[:len(cmd)-1]
		if cmd[0] == ".." {
			if len(cur) > 0 {
				cur = cur[:len(cur)-1]
			}
		} else if cmd[0] != "." {
			for _, l := range cmd {
				cur = append(cur, l)
			}
		}
	}
	fmt.Println(cur)
	return len(cur)
}
