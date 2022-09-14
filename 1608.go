package Question_by_day

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] >= (n - i) {
			if i == 0 || nums[i-1] < n-i {
				return n - i
			}
		}
	}
	return -1
}
