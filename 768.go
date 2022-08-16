package Question_by_day

import "sort"

func maxChunksToSorted(arr []int) int {
	sortedArr := make([]int, len(arr))
	for id, v := range arr {
		sortedArr[id] = v
	}
	sort.Ints(sortedArr)
	del := make(map[int]int)
	res := 0
	for i := 0; i < len(arr); i++ {
		del[arr[i]] += 1
		if del[arr[i]] == 0 {
			delete(del, arr[i])
		}
		del[sortedArr[i]] -= 1
		if del[sortedArr[i]] == 0 {
			delete(del, sortedArr[i])
		}
		if len(del) == 0 {
			res += 1
		}
	}

	return res
}
