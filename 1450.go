package Question_by_day

func busyStudent(startTime []int, endTime []int, queryTime int) (res int) {
	n := len(startTime)
	for i := 0; i < n; i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			res += 1
		}
	}
	return
}
