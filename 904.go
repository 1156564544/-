package Question_by_day

func totalFruit(fruits []int) int {
	var find func(fruits []int, start, end int) int
	find = func(fruits []int, start, end int) int {
		mid := (start + end) / 2
		first, second := fruits[mid], -1
		r1 := 1
		for j := mid + 1; j <= end; j++ {
			if fruits[j] == first {
				r1 += 1
			} else if second == -1 || second == fruits[j] {
				second = fruits[j]
				r1 += 1
			} else {
				break
			}
		}
		for j := mid - 1; j >= start; j-- {
			if fruits[j] == first {
				r1 += 1
			} else if second == -1 || second == fruits[j] {
				second = fruits[j]
				r1 += 1
			} else {
				break
			}
		}
		r2 := 1
		first, second = fruits[mid], -1
		for j := mid - 1; j >= start; j-- {
			if fruits[j] == first {
				r2 += 1
			} else if second == -1 || second == fruits[j] {
				second = fruits[j]
				r2 += 1
			} else {
				break
			}
		}
		for j := mid + 1; j <= end; j++ {
			if fruits[j] == first {
				r2 += 1
			} else if second == -1 || second == fruits[j] {
				second = fruits[j]
				r2 += 1
			} else {
				break
			}
		}
		return max(r1, r2)
	}
	var binary func(fruits []int, start, end int) int
	binary = func(fruits []int, start, end int) int {
		if end-start <= 1 {
			return end - start + 1
		}
		mid := (start + end) / 2
		r1 := binary(fruits, start, mid-1)
		r2 := binary(fruits, mid+1, end)
		res := max(r1, r2)
		res = max(res, find(fruits, start, end))
		return res
	}

	return binary(fruits, 0, len(fruits)-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
