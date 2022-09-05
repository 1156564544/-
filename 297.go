package Question_by_day

import "strconv"

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	if root.Left == nil && root.Right == nil {
		return strconv.Itoa(root.Val)
	}
	res := []byte{}
	if root.Left != nil {
		rl_ := this.serialize(root.Left)
		res = append(res, '(')
		res = append(res, []byte(rl_)...)
		res = append(res, ')')
	}
	res = append(res, []byte(strconv.Itoa(root.Val))...)
	if root.Right != nil {
		rr_ := this.serialize(root.Right)
		res = append(res, '(')
		res = append(res, []byte(rr_)...)
		res = append(res, ')')
	}
	return string(res)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	ss := []byte(data)
	if len(ss) == 0 {
		return nil
	}
	root := new(TreeNode)
	idx2 := 0
	if ss[0] != '(' {
		root.Left = nil
		idx2 = 0
		for idx2 < len(ss) {
			if ss[idx2] == '(' {
				break
			}
			idx2 += 1
		}
		root.Val, _ = strconv.Atoi(string(ss[:idx2]))
	} else {
		idx := 1
		num := 1
		for idx < len(ss) {
			if ss[idx] == '(' {
				num += 1
			}
			if ss[idx] == ')' {
				num -= 1
				if num == 0 {
					break
				}
			}
			idx += 1
		}
		root.Left = this.deserialize(string(ss[1:idx]))
		idx2 = idx + 1
		for idx2 < len(ss) {
			if ss[idx2] == '(' {
				break
			}
			idx2 += 1
		}
		root.Val, _ = strconv.Atoi(string(ss[idx+1 : idx2]))
	}
	if idx2 == len(ss) {
		root.Right = nil
	} else {
		root.Right = this.deserialize(string(ss[idx2+1 : len(ss)-1]))
	}
	return root
}
