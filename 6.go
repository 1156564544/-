package Question_by_day

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ss := []byte(s)
	n := len(ss)
	res := []byte{}
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			for j := i; j < n; j = j + 2*numRows - 2 {
				res = append(res, ss[j])
			}
		} else {
			flag := true
			j := i
			k := 2*numRows - i - 2
			for j < n || k < n {
				if flag {
					res = append(res, ss[j])
					flag = false
					j = j + 2*numRows - 2
				} else {
					res = append(res, ss[k])
					flag = true
					k = k + 2*numRows - 2
				}
			}
		}
	}
	return string(res)
}
