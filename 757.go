package Question_by_day

// 用贪心算法来求解，首先对intervals进行排序，每次添加数据时添加interval最左边的数

import "sort"

type intervalType [][]int

func (i intervalType) Len() int {
	return len(i)
}

func (i intervalType) Less(x, y int) bool {
	if i[x][0] < i[y][0] {
		return true
	}
	if i[x][0] == i[y][0] {
		if i[x][1] >= i[y][1] {
			return true
		}
	}
	return false
}

func (i intervalType) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func intersectionSizeTwo(intervals [][]int) int {
	sort.Sort(intervalType(intervals))
	n := len(intervals)
	set := []int{}
	set = append(set, intervals[n-1][0]+1)
	set = append(set, intervals[n-1][0])
	lists := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		lists[i] = make([]int, 0)
		if intervals[i][0] <= set[0] && intervals[i][1] >= set[0] {
			lists[i] = append(lists[i], set[0])
		}
		if intervals[i][0] <= set[1] && intervals[i][1] >= set[1] {
			lists[i] = append(lists[i], set[1])
		}
	}
	for i := n - 2; i >= 0; i-- {
		for len(lists[i]) < 2 {
			for j := intervals[i][0]; j <= intervals[i][1] && len(lists[i]) < 2; j++ {
				if len(lists[i]) == 0 || lists[i][0] != j {
					lists[i] = append(lists[i], j)
					for k := i - 1; k >= 0; k-- {
						if len(lists[k]) == 2 {
							continue
						}
						if j >= intervals[k][0] && j <= intervals[k][1] {
							lists[k] = append(lists[k], j)
						}
					}
					set = append(set, j)
				}
			}
		}
	}

	return len(set)
}

func isOverlapped(interval []int, set []int) int {
	sum := 0
	idx := len(set) - 1
	for i := interval[0]; i <= interval[1]; i++ {
		if i == set[idx] {
			sum += 1
			idx -= 1
		}
		if sum == 2 {
			return 2
		}
	}
	return sum
}
