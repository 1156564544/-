package Question_by_day

func shuffle(nums []int, n int) []int {
	tmp := make([]int, n)
	copy(tmp, nums[n:])
	for i := n - 1; i > 0; i-- {
		nums[2*i] = nums[i]
	}
	for i := 0; i < n; i++ {
		nums[2*i+1] = tmp[i]
	}
	return nums
}
