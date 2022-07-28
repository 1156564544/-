package Question_by_day

import "sort"

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	help := make([]int, len(arr))
	for i, v := range arr {
		help[i] = v
	}
	sort.Ints(help)

	m := map[int]int{}
	idx := 1
	m[help[0]] = idx
	for i := 1; i < len(help); i++ {
		if help[i] == help[i-1] {
			m[help[i]] = m[help[i-1]]
		} else {
			idx += 1
			m[help[i]] = idx
		}
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = m[arr[i]]
	}
	return arr
}
