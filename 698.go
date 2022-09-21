package Question_by_day

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	total := sum / k
	n := len(nums)
	sort.Ints(nums)
	if nums[n-1] > total {
		return false
	}
	state := make([]bool, 1<<n)
	totalSum := make([]int, 1<<n)
	state[0] = true
	for i, v := range state {
		if !v {
			continue
		}
		for j, k := range nums {
			if totalSum[i]+k > total {
				break
			}
			if (i>>j)&1 == 0 {
				next := i | (1 << j)
				if !state[next] {
					totalSum[next] = (totalSum[i] + nums[j]) % total
					state[next] = true

				}
			}
		}
	}
	return state[1<<n-1]
}
