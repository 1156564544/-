package Question_by_day

import "sort"

type array [][]int

func (a array) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a array) Len() int {
	return len(a)
}

func (a array) Less(i, j int) bool {
	if a[i][0] < a[j][0] {
		return true
	} else if a[i][0] == a[j][0] {
		return a[i][1] < a[j][1]
	}
	return false
}

func findLongestChain(pairs [][]int) int {
	a := array(pairs)
	sort.Sort(a)
	ppairs := make([][]int, 1)
	ppairs[0] = a[0]
	for i := 1; i < len(a); i++ {
		if a[i][0] == a[i-1][0] {
			continue
		}
		ppairs = append(ppairs, a[i])
	}
	n := len(ppairs)
	count := make([]int, n)
	count[0] = 1
	for j := 1; j < n; j++ {
		var i int
		for i = j - 1; i >= 0; i-- {
			if ppairs[i][1] < ppairs[j][0] {
				break
			}
		}
		if i < 0 {
			count[j] = 1
		} else {
			count[j] = count[i] + 1
		}
	}
	return count[n-1]
}
