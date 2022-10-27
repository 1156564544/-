package Question_by_day

func isIdealPermutation(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return true
	}
	max1 := max(nums[0], nums[1])
	max2 := nums[0] + nums[1] - max1
	for i := 2; i < n; i++ {
		if nums[i] > nums[i-1] {
			if nums[i] < max1 {
				return false
			} else {
				max1, max2 = nums[i], max1
			}
		} else {
			if nums[i] < max2 {
				return false
			} else {
				max2 = nums[i]
			}
		}
	}
	return true
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
