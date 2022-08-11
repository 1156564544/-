package Question_by_day

func minStartValue(nums []int) int {
	s := nums[0]
	minS := nums[0]
	for i := 1; i < len(nums); i++ {
		s += nums[i]
		if s < minS {
			minS = s
		}
	}
	if minS > 0 {
		return 1
	} else {
		return 1 - minS
	}
}
