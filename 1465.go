package Question_by_day

import "sort"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	horizontalCuts = append(horizontalCuts, h)
	verticalCuts = append(verticalCuts, w)
	h_m, w_m := horizontalCuts[0], verticalCuts[0]
	for i := len(horizontalCuts) - 1; i > 0; i-- {
		horizontalCuts[i] = horizontalCuts[i] - horizontalCuts[i-1]
		if h_m < horizontalCuts[i] {
			h_m = horizontalCuts[i]
		}
	}
	for i := len(verticalCuts) - 1; i > 0; i-- {
		verticalCuts[i] = verticalCuts[i] - verticalCuts[i-1]
		if w_m < verticalCuts[i] {
			w_m = verticalCuts[i]
		}
	}
	return (h_m * w_m) % (1000000007)
}
