package Question_by_day

func maxSum(nums1 []int, nums2 []int) int {
	var res int64 = 0
	var a, b int64
	i, j := 0, 0
	n, m := len(nums1), len(nums2)
	for i < n && j < m {
		if nums1[i] == nums2[j] {
			a += int64(nums1[i])
			b += int64(nums2[j])
			res += max(a, b)
			a, b = 0, 0
			i += 1
			j += 1
		} else if nums1[i] < nums2[j] {

			a += int64(nums1[i])
			i += 1
		} else {

			b += int64(nums2[j])
			j += 1
		}
	}
	for i < n {
		a += int64(nums1[i])
		i += 1
	}
	for j < m {
		b += int64(nums2[j])
		j += 1
	}
	res += max(a, b)
	return int(res % 1000000007)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
