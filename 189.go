package Question_by_day

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	temp := 0
	count := 0
	for i := n - 1; i > n-1-k; i-- {
		temp = nums[i]
		var j int
		for j = (i - k + n) % n; j != i; j = (j - k + n) % n {
			nums[(j+k)%n] = nums[j]
			count += 1
		}
		nums[(i+k)%n] = temp
		count += 1
		if count == n {
			break
		}
	}
}
