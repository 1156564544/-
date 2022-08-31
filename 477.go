package Question_by_day

func totalHammingDistance(nums []int) int {
	bits := make([]int, 32)
	n := len(nums)
	var help func(num int)
	help = func(num int) {
		mask := 1
		for i := 0; i < 32; i++ {
			if num&mask != 0 {
				bits[i] += 1
			}
			mask = mask << 1
		}
	}
	for i := 0; i < n; i++ {
		help(nums[i])
	}

	res := 0
	for i := 0; i < 32; i++ {
		res += (n - bits[i]) * bits[i]
	}
	return res
}
