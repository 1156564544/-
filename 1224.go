package Question_by_day

func maxEqualFreq(nums []int) int {
	fre := make(map[int]int)
	maxFre := 0
	m := make(map[int]int)
	res := 0
	for i, n := range nums {
		if _, ok := fre[n]; ok {
			m[fre[n]] -= 1
			if m[fre[n]] == 0 {
				delete(m, fre[n])
			}
		}
		fre[n] += 1
		if maxFre < fre[n] {
			maxFre = fre[n]
		}
		m[fre[n]] += 1

		if maxFre == 1 {
			res = i + 1
			continue
		}
		if m[1] == 1 && m[maxFre] == len(fre)-1 {
			res = i + 1
			continue
		}
		if m[maxFre] == 1 && m[maxFre-1] == len(fre)-1 {
			res = i + 1
		}
	}
	return res
}
