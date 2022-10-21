package Question_by_day

import "sort"

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	idx1 := make([]int, n)
	for i := 1; i < n; i++ {
		idx1[i] = i
	}
	sort.Slice(idx1, func(i, j int) bool {
		return position[idx1[i]] > position[idx1[j]]
	})
	res := 0
	for i := 0; i < n; i++ {
		time := float64(target-position[idx1[i]]) / float64(speed[idx1[i]])
		res += 1
		var j int
		for j = i + 1; j < n; j++ {
			time2 := float64(target-position[idx1[j]]) / float64(speed[idx1[j]])
			if time2 > time {
				i = j - 1
				break
			}
		}
		if j == n {
			break
		}
	}

	return res
}
