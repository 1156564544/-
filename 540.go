package Question_by_day

func singleNonDuplicate(nums []int) int {
	return help(0, len(nums)-1, nums)
}

func help(start, end int, nums []int) int {
	if start == end {
		return nums[start]
	}
	mid := (start + end) / 2
	if nums[mid] == nums[mid-1] {
		if (mid-start)%2 == 0 {
			return help(start, mid-2, nums)
		} else {
			return help(mid+1, end, nums)
		}
	} else if nums[mid] == nums[mid+1] {
		if (mid-start)%2 != 0 {
			return help(start, mid-1, nums)
		} else {
			return help(mid, end, nums)
		}
	} else {
		return nums[mid]
	}
}
