package Question_by_day

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	ans := 1
	sum := root.Val
	queue := []*TreeNode{root}
	floor := 0
	for len(queue) > 0 {
		floor += 1
		s := 0
		q := []*TreeNode{}
		for _, node := range queue {
			s += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if sum < s {
			sum = s
			ans = floor
		}
		queue = q
	}
	return ans
}
