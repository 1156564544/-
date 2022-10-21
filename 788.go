package Question_by_day

import "strconv"

func rotatedDigits(n int) int {
	count := []int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}
	numdiff := []int{1, 2, 3, 3, 3, 4, 5, 5, 6, 7}
	numnodiff := []int{0, 0, 1, 1, 1, 2, 3, 3, 3, 4}
	mem := [5][2][2]int{}
	for i := 0; i < 5; i++ {
		mem[i] = [2][2]int{{-1, -1}, {-1, -1}}
	}
	var dfs func(int, bool, bool) int
	dfs = func(pos int, bound bool, diff bool) int {
		digital := strconv.Itoa(n)
		ptr := &mem[pos][bool2int(bound)][bool2int(diff)]
		if *ptr != -1 {
			return *ptr
		}
		if pos == len(digital)-1 {
			if !bound {
				if diff {
					*ptr = 7
				} else {
					*ptr = 4
				}
				return *ptr
			}
			if diff {
				*ptr = numdiff[digital[pos]-'0']
			} else {
				*ptr = numnodiff[digital[pos]-'0']
			}
			return *ptr
		}
		lim := 9
		if bound {
			lim = int(digital[pos] - '0')
		}
		*ptr = 0
		for i := 0; i <= lim; i++ {
			if count[i] == -1 {
				continue
			}
			*ptr += dfs(pos+1, bound && i == int(digital[pos]-'0'), diff || count[i] == 1)

		}
		return *ptr
	}

	return dfs(0, true, false)
}

func bool2int(f bool) int {
	if f {
		return 1
	}
	return 0
}
