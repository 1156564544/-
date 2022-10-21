package Question_by_day

func buildArray(target []int, n int) []string {
	start := 1
	res := make([]string, 0)
	for _, v := range target {
		for v > start {
			res = append(res, "Push", "Pop")
			start += 1
		}
		res = append(res, "Push")
		start += 1
	}
	return res
}
