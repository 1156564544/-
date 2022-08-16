package Question_by_day

func maxChunksToSorted(arr []int) int {
	res := 0
	for i := 0; i < len(arr); {
		end := arr[i]
		res += 1
		i += 1
		for ; i <= end; i++ {
			if arr[i] > end {
				end = arr[i]
			}
		}
	}
	return res
}
