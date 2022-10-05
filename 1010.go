package Question_by_day

import (
	"fmt"
	"sort"
)

func numPairsDivisibleBy60(time []int) int {
	for i, v := range time {
		time[i] = v % 60
	}
	sort.Ints(time)
	a1 := []int{time[0]}
	a2 := []int{1}
	j := 0
	for i := 1; i < len(time); i++ {
		if time[i] == time[i-1] {
			a2[j] += 1
		} else {
			a1 = append(a1, time[i])
			a2 = append(a2, 1)
			j += 1
		}
	}
	fmt.Println(a1, a2)
	res := 0
	i := 0
	if a1[i] == 0 {
		res += leijia(a2[i])
		i += 1
	}
	j = len(a1) - 1
	for i < j {
		if a1[i]+a1[j] == 60 {
			res += a2[i] * a2[j]
			i += 1
			j -= 1
		} else if a1[i]+a1[j] < 60 {
			i += 1
		} else {
			j -= 1
		}
	}
	if i == j && a1[i] == 30 {
		res += leijia(a2[i])
	}

	return res
}

func leijia(n int) int {
	n = n - 1
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
