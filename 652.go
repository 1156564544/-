package Question_by_day

import (
	"fmt"
	"strconv"
)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := []*TreeNode{}
	m := make(map[string]int)
	var dfs func(*TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			return ""
		}
		ss := fmt.Sprintf("(%s)%s(%s)", dfs(root.Left), strconv.Itoa(root.Val), dfs(root.Right))
		m[ss] += 1
		if m[ss] == 2 {
			res = append(res, root)
		}
		return ss
	}
	dfs(root)
	return res
}
