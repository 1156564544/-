package Question_by_day

func mirrorReflection(p int, q int) int {
	gbs := help(p, q)
	a := gbs / p
	b := gbs / q

	if a%2 == 0 {
		return 0
	}
	if b%2 == 0 {
		return 2
	}
	return 1
}

func help(x, y int) int {
	var i int
	for i = x; i <= x*y; i += x {
		if i%y == 0 {
			break
		}
	}
	return i
}
