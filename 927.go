package Question_by_day

func threeEqualParts(arr []int) []int {
	sum := 0
	for _, v := range arr {
		if v == 1 {
			sum += 1
		}
	}
	if sum == 0 {
		return []int{0, 2}
	}
	if sum%3 != 0 {
		return []int{-1, -1}
	}
	partial := sum / 3
	index := make([]int, sum)
	idx := 0
	for i, v := range arr {
		if v == 1 {
			index[idx] = i
			idx += 1
		}
	}
	c, d := index[2*partial-1], index[2*partial]
	i := index[0] + len(arr) - d - 1
	j := d
	if less(arr[:i+1], arr[j:]) != 0 {
		return []int{-1, -1}
	}
	for j > c {
		x := less(arr[:i+1], arr[i+1:j])
		if x < 0 {
			j -= 1
			continue
		}
		if x > 0 {
			break
		}
		return []int{i, j}
	}

	return []int{-1, -1}
}

func less(a []int, b []int) int {
	n, m := len(a), len(b)
	i, j := 0, 0
	for i < n && j < m {
		for i < n && a[i] == 0 {
			i += 1
		}
		for j < m && b[j] == 0 {
			j += 1
		}
		if i == n && j == m {
			return 0
		}
		if i == n && j < m {
			return -1
		}
		if i < n && j == m {
			return 1
		}
		if n-i > m-j {
			return 1
		}
		if n-i < m-j {
			return -1
		}
		i += 1
		j += 1
	}
	return 0
}
