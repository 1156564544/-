package Question_by_day

func replaceSpace(s string) string {
	s_ := []byte(s)
	index := []int{}
	for i, ch := range s_ {
		if ch == ' ' {
			index = append(index, i)
		}
	}
	res := make([]byte, len(s_)+len(index)*2)
	last := 0
	for i, idx := range index {
		for j := last; j < idx; j++ {
			res[j+i*2] = s_[j]
		}
		res[idx+i*2] = '%'
		res[idx+i*2+1] = '2'
		res[idx+i*2+2] = '0'
		last = idx + 1
	}
	for j := last; j < len(s_); j++ {
		res[j+len(index)*2] = s_[j]
	}

	return string(res)
}
