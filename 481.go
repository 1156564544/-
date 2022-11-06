package Question_by_day

func magicalString(n int) int {
	if n <= 3 {
		return 1
	}
	s := []int{1, 2, 2}
	i := 2
	num := 2
	res := 1
	for len(s) < n {
		num = num%2 + 1
		for j := 0; j < s[i] && len(s) < n; j++ {
			s = append(s, num)
			res += num % 2
		}
		i += 1
	}
	return res
}
