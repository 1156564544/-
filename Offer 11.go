package Question_by_day

func minArray(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	if len(numbers) == 2 {
		return min(numbers[0], numbers[1])
	}
	n := len(numbers)
	mid := n / 2
	if numbers[0] == numbers[mid] {
		return min(minArray(numbers[:mid+1]), minArray(numbers[mid+1:]))
	}
	if numbers[0] > numbers[mid] {
		return minArray(numbers[:mid+1])
	} else {
		return min(numbers[0], minArray(numbers[mid+1:]))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
