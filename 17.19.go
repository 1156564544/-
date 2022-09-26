package Question_by_day

func missingTwo(nums []int) []int {
	N := len(nums) + 2
	a := N / 2
	c1 := 0
	sum1, sum2 := 0, 0
	for _, v := range nums {
		if v%2 == 0 {
			c1 += 1
			sum1 += v
		} else {
			sum2 += v
		}
	}
	sum3, sum4 := 0, 0
	for i := 1; i <= N; i++ {
		if i%2 == 0 {
			sum3 += i
		} else {
			sum4 += i
		}
	}
	if c1 == a-1 {
		return []int{sum3 - sum1, sum4 - sum2}
	} else if c1 == a-2 {
		mean := (sum3 - sum1) / 2
		sum5 := 0
		for _, v := range nums {
			if v%2 == 0 && v <= mean {
				sum5 += v
			}
		}
		for i := 2; i <= mean; i += 2 {
			sum5 -= i
		}
		return []int{-sum5, sum3 - sum1 + sum5}
	} else {
		mean := (sum4 - sum2) / 2
		sum5 := 0
		for _, v := range nums {
			if v%2 == 1 && v <= mean {
				sum5 += v
			}
		}
		for i := 1; i <= mean; i += 2 {
			sum5 -= i
		}
		return []int{-sum5, sum4 - sum2 + sum5}
	}
}
