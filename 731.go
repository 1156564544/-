package Question_by_day

// 通过线段树来求解本题
type pair struct {
	n    int
	lazy int
}

type MyCalendarTwo struct {
	tree map[int]pair
}

func Constructor() MyCalendarTwo {
	var c MyCalendarTwo
	c.tree = make(map[int]pair)
	return c
}

func (this *MyCalendarTwo) update(start, end, l, r, val, idx int) {
	if start > r || end < l {
		return
	}
	if start <= l && end >= r {
		node := this.tree[idx]
		node.n += val
		node.lazy += val
		this.tree[idx] = node
		return
	}
	mid := (l + r) / 2
	this.update(start, end, l, mid, val, idx*2)
	this.update(start, end, mid+1, r, val, idx*2+1)
	node := this.tree[idx]
	left := this.tree[2*idx]
	right := this.tree[2*idx+1]
	node.n = max(left.n, right.n) + node.lazy
	this.tree[idx] = node
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	this.update(start, end-1, 0, 1e9, 1, 1)
	if (this.tree[1]).n > 2 {
		this.update(start, end-1, 0, 1e9, -1, 1)
		return false
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
