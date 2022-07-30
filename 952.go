package Question_by_day

// 用并查集来做

type set struct {
	parent, rank []int
}

func NewSet(n int) set {
	parent := make([]int, n+1)
	for i := 0; i <= n; i++ {
		parent[i] = i
	}
	return set{parent, make([]int, n+1)}
}

func (s set) find(x int) int {
	if s.parent[x] != x {
		return s.find(s.parent[x])
	}
	return x
}

func (s set) merge(x, y int) {
	xp, yp := s.find(x), s.find(y)
	if xp == yp {
		return
	}
	if s.rank[xp] >= s.rank[yp] {
		s.parent[yp] = xp
		if s.rank[xp] == s.rank[yp] {
			s.rank[xp] += 1
		}
	} else {
		s.parent[xp] = yp
	}
}

func largestComponentSize(nums []int) int {
	m := 0
	for _, i := range nums {
		m = max(m, i)
	}
	s := NewSet(m)
	for _, num := range nums {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				s.merge(num, i)
				s.merge(num, num/i)
			}
		}
	}
	res := 0
	cnt := make([]int, m+1)
	for _, num := range nums {
		cnt[s.find(num)] += 1
		res = max(res, cnt[s.find(num)])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
