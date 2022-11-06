package Question_by_day

func reachNumber(target int) int {
	target = max(target, -target)
	res := 1
	length := 1
	for length < target {
		res += 1
		length += res
	}
	for length > target && (length-target)%2 != 0 {
		res += 1
		length += res
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
