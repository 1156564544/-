package Question_by_day

import "sort"

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)
	n, m := len(buses), len(passengers)
	j := 0
	cur := 0
	for i := 0; i < n-1; i++ {
		cur = 0
		for j < m && cur < capacity {
			if passengers[j] <= buses[i] {
				cur += 1
				j += 1
			} else {
				break
			}
		}
	}
	for cur = 0; cur < capacity-1 && j < m; cur++ {
		if passengers[j] <= buses[n-1] {
			j++
		} else {
			break
		}
	}
	var last int
	if j == m {
		last = buses[n-1]
	} else {
		last = passengers[j] - 1
	}

	if last > buses[n-1] {
		last = buses[n-1]
	}
	j -= 1
	for j > 0 {
		if last > passengers[j] {
			break
		} else if last == passengers[j] {
			last = last - 1
			j -= 1
		} else {
			j -= 1
		}
	}
	if j == 0 && last == passengers[j] {
		return last - 1
	}
	return last
}
