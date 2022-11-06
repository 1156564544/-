package Question_by_day

func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	ans := make([]int, 0)
	if time == 0 {
		for i := 0; i < n; i++ {
			ans = append(ans, i)
		}
		return ans
	}
	left := make([]int, n)
	right := make([]int, n)
	left[0] = 0
	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 0
		}
	}
	right[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			right[i] = right[i+1] + 1
			if right[i] >= time && left[i] >= time {
				ans = append(ans, i)
			}
		}
	}
	return ans
}
