package Question_by_day

func numberOfPairs(nums1 []int, nums2 []int, diff int) int64 {
	nums3 := make([]int, len(nums1))
	for i, v := range nums1 {
		nums3[i] = v - nums2[i]
	}
	return help(nums3, diff)
}

func help(nums []int, diff int) int64 {
	n := len(nums)
	if n < 2 {
		return 0
	}
	if n == 2 {
		var res int64 = 0
		if nums[0] <= nums[1]+diff {
			res = 1
		}
		if nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return res
	}
	s1 := help(nums[:n/2], diff)
	s2 := help(nums[n/2:], diff)
	res := s1 + s2
	arr := make([]int, n)
	i, j := 0, n/2
	for i < n/2 && j < n {
		if nums[i] <= nums[j]+diff {
			res += int64(n - j)
			i += 1
		} else {
			j += 1
		}
	}
	i, j = 0, n/2
	for i < n/2 && j < n {
		if nums[i] < nums[j] {
			arr[i+j-n/2] = nums[i]
			i += 1
		} else {
			arr[i+j-n/2] = nums[j]
			j += 1
		}
	}
	if i == n/2 {
		for j < n {
			arr[j] = nums[j]
			j += 1
		}
	} else {
		for i < n/2 {
			arr[i+n-n/2] = nums[i]
			i += 1
		}
	}
	copy(nums, arr)
	return res
}
