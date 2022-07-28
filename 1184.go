package Question_by_day

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	x, y := 0, 0
	for i := start; i != destination; i = (i + 1) % len(distance) {
		x += distance[i]
	}
	for i := destination; i != start; i = (i + 1) % len(distance) {
		y += distance[i]
	}

	if x < y {
		return x
	}
	return y
}
