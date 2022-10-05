package Question_by_day

func canTransform(start string, end string) bool {

	if len(start) != len(end) {
		return false
	}
	n := len(start)
	i, j := 0, 0
	for i < n && j < n {
		for i < n && start[i] == 'X' {
			i += 1
		}
		for j < n && end[j] == 'X' {
			j += 1
		}
		if i == n && j != n {
			return false
		}
		if i != n && j == n {
			return false
		}
		if i == n && j == n {
			return true
		}
		if start[i] != end[j] {
			return false
		}
		if start[i] == 'L' {
			if i < j {
				return false
			}
		} else {
			if i > j {
				return false
			}
		}
		i += 1
		j += 1
	}
	if i != n {
		for i < n {
			if start[i] != 'X' {
				return false
			}
			i += 1
		}
	}
	if j != n {
		for j < n {
			if end[j] != 'X' {
				return false
			}
			j += 1
		}
	}
	return true
}
