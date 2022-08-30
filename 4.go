package Question_by_day

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n%2 == 1 {
		return float64(find(nums1, nums2, n/2+1))
	} else {
		return (float64(find(nums1, nums2, n/2)) + float64(find(nums1, nums2, n/2+1))) / 2.0
	}
}

func find(nums1 []int, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if len(nums2) == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}
	idx1 := min(k/2-1, len(nums1)-1)
	idx2 := min(k/2-1, len(nums2)-1)
	if nums1[idx1] <= nums2[idx2] {
		return find(nums1[idx1+1:], nums2, k-idx1-1)
	} else {
		return find(nums1, nums2[idx2+1:], k-idx2-1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
