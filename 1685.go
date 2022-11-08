package Question_by_day

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 1; i < n; i++ {
		res[0] += nums[i] - nums[0]
	}
	for i := 1; i < n; i++ {
		diff := nums[i] - nums[i-1]
		res[i] = res[i-1] + diff*i - diff*(n-i)
	}
	return res
}
