package Question_by_day

import "sort"

func maxOperations(nums []int, k int) int {
	ans := 0
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] < k {
			i += 1
		} else if nums[i]+nums[j] > k {
			j -= 1
		} else {
			i += 1
			j -= 1
			ans += 1
		}
	}
	return ans
}
