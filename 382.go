package Question_by_day

import "math/rand"

type Solution struct {
	node *ListNode
	num  int
}

func Constructor(head *ListNode) Solution {
	num := 0
	for cur := head; cur != nil; cur = cur.Next {
		num += 1
	}
	return Solution{head, num}
}

func (this *Solution) GetRandom() int {
	step := rand.Intn(this.num)
	cur := this.node
	for i := 0; i < step; i++ {
		cur = cur.Next
	}
	return cur.Val
}
