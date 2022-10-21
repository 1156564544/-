package Question_by_day

func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	c1, c2 := 0, 1
	for i := 1; i < n; i++ {
		c3, c4 := -1, -1
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			c3 = c1
			c4 = c2 + 1
		}
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			if c3 != -1 {
				c3 = min(c3, c2)
			} else {
				c3 = c2
			}
			if c4 != -1 {
				c4 = min(c4, c1+1)
			} else {
				c4 = c1 + 1
			}
		}
		c1, c2 = c3, c4
	}

	return min(c1, c2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
