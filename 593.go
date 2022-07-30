package Question_by_day

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	// fmt.Println(vertical(p1,p2,p3,p4))
	// fmt.Println(equal(p1,p2,p3,p4))
	// fmt.Println(zd(p1,p2,p3,p4))
	f1 := vertical(p1, p2, p3, p4) && equal(p1, p2, p3, p4) && zd(p1, p2, p3, p4)
	f2 := vertical(p1, p3, p2, p4) && equal(p1, p3, p2, p4) && zd(p1, p3, p2, p4)
	f3 := vertical(p1, p4, p2, p3) && equal(p1, p4, p2, p3) && zd(p1, p4, p2, p3)
	return f1 || f2 || f3
}

func vertical(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	y := (p2[1] - p1[1]) * (p4[1] - p3[1])
	x := (p2[0] - p1[0]) * (p4[0] - p3[0])
	return x == -y
}

func equal(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	d1 := (p2[1]-p1[1])*(p2[1]-p1[1]) + (p2[0]-p1[0])*(p2[0]-p1[0])
	d2 := (p4[1]-p3[1])*(p4[1]-p3[1]) + (p4[0]-p3[0])*(p4[0]-p3[0])
	if d1 == 0 {
		return false
	}
	return d1 == d2
}

func zd(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	x1 := (p1[0] + p2[0]) / 2
	y1 := (p1[1] + p2[1]) / 2
	x2 := (p3[0] + p4[0]) / 2
	y2 := (p3[1] + p4[1]) / 2
	// fmt.Println(x1,x2,y1,y2)
	return x1 == x2 && y1 == y2
}
