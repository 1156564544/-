package Question_by_day

import "sort"

func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums2)
	idx1 := make([]int, n)
	idx2 := make([]int, n)
	for i := 1; i < n; i++ {
		idx1[i] = i
		idx2[i] = i
	}

	sort.Slice(idx1, func(i, j int) bool {
		return nums1[idx1[i]] < nums1[idx1[j]]
	})
	sort.Slice(idx2, func(i, j int) bool {
		return nums2[idx2[i]] < nums2[idx2[j]]
	})
	i, j, k := 0, 0, n-1
	res := make([]int, n)
	for ; i < n && j < n; i++ {
		if nums1[idx1[i]] <= nums2[idx2[j]] {
			res[idx2[k]] = nums1[idx1[i]]
			k -= 1
		} else {
			res[idx2[j]] = nums1[idx1[i]]
			j += 1
		}
	}
	return res
}
