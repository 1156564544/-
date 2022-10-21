package Question_by_day

import "sort"

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	sort.Ints(arr2)
	arr3 := make([]int, 1)
	arr3[0] = arr2[0]
	for i := 1; i < len(arr2); i++ {
		if arr2[i] == arr2[i-1] {
			continue
		}
		arr3 = append(arr3, arr2[i])
	}
	var find func(arr []int, start, end, v int) int
	find = func(arr []int, start, end, v int) int {
		if start > end {
			return -1
		}
		if start == end {
			if arr[start] < v {
				return start
			} else {
				return -1
			}
		}
		mid := (start + end) / 2
		if arr[mid] < v {
			res := find(arr, mid+1, end, v)
			if res != -1 {
				return res
			} else {
				return mid
			}
		}
		return find(arr, start, mid-1, v)
	}
	arr1 = append(arr1, 1<<30)
	n := len(arr1)
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		v := arr1[i]
		if v > arr1[i-1] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = 1 << 30
		}
		j := find(arr3, 0, len(arr3)-1, v)
		for k := 1; k <= min(i, j+1); k++ {
			if i == k {
				dp[i] = min(dp[i], k)
			} else if arr1[i-k-1] < arr3[j-k+1] {
				dp[i] = min(dp[i], dp[i-k-1]+k)
			}
		}
		if arr1[i-1] < arr1[i] {
			dp[i] = min(dp[i], dp[i-1])
		}
	}
	if dp[n-1] == 1<<30 {
		return -1
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
