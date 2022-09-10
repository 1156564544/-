package Question_by_day

import (
	"fmt"
	"sort"
)

func maximumImportance(n int, roads [][]int) int64 {
	dushu := make([]int, n)
	for _, r := range roads {
		dushu[r[0]] += 1
		dushu[r[1]] += 1
	}
	sort.Ints(dushu)
	fmt.Println(dushu)
	var res int64
	count := 1
	for _, d := range dushu {
		res += int64(d * count)
		count += 1
	}
	return res
}
