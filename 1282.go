package Question_by_day

func groupThePeople(groupSizes []int) [][]int {
	m := make(map[int][]int)
	for id, size := range groupSizes {
		if _, ok := m[size]; !ok {
			m[size] = []int{id}
		} else {
			m[size] = append(m[size], id)
		}
	}
	res := make([][]int, 0)
	for k, v := range m {
		for i := 0; i < len(v); i += k {
			cur := make([]int, k)
			for j := i; j < i+k; j++ {
				cur[j-i] = v[j]
			}
			res = append(res, cur)
		}
	}

	return res
}
