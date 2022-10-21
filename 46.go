package Question_by_day

func permute(nums []int) (res [][]int) {
	n := len(nums)
	var dfs func(del []bool, cur []int)
	dfs = func(del []bool, cur []int) {
		if len(cur) == n {
			// 切片处理
			res = append(res, append([]int{}, cur...))
			return
		}
		for i := 0; i < n; i++ {
			if del[i] {
				continue
			}
			del[i] = true
			cur = append(cur, nums[i])
			dfs(del, cur)
			del[i] = false
			cur = cur[:len(cur)-1]
		}
	}
	delete := make([]bool, n)
	dfs(delete, []int{})
	return
}
