package Question_by_day

func canBeEqual(target []int, arr []int) bool {
	m := make(map[int]int)
	for _, v := range target {
		m[v] += 1
	}
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			return false
		}
		m[v] -= 1
		if m[v] == 0 {
			delete(m, v)
		}
	}
	return len(m) == 0
}
