package Question_by_day

func sumSubarrayMins(arr []int) int {
	var ans int64
	for i := 0; i < len(arr); i++ {
		var j, l int
		for j = i - 1; j >= 0; j-- {
			if arr[j] < arr[i] {
				break
			}
		}
		for l = i + 1; l < len(arr); l++ {
			if arr[l] <= arr[i] {
				break
			}
		}
		ans += int64((i - j) * (l - i) * arr[i])
	}
	return int(ans % 1000000007)
}
