package Question_by_day

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	n := len(quality)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = i
	}
	sort.Slice(h, func(i, j int) bool {
		a, b := h[i], h[j]
		return wage[a]*quality[b] < wage[b]*quality[a]
	})
	totalq := 0
	q := hp{}
	for i := 0; i < k-1; i++ {
		totalq += quality[h[i]]
		heap.Push(&q, quality[h[i]])
	}
	ans := 1e9
	for i := k - 1; i < n; i++ {
		totalq += quality[h[i]]
		heap.Push(&q, quality[h[i]])
		ans = math.Min(ans, float64(totalq)*float64(wage[h[i]])/float64(quality[h[i]]))
		totalq -= heap.Pop(&q).(int)
	}
	return ans
}

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	h.IntSlice = a[:len(a)-1]
	return a[len(a)-1]
}
