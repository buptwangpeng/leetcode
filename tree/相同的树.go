package tree

import "leetcode/utils"

// leetcode easy 100

func isSameTree(p *utils.TreeNode, q *utils.TreeNode) bool {
	// 处理root
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if p.Value != q.Value {
		return false
	}

	// 递归
	return isSameTree(p.LeftNode, q.LeftNode) && isSameTree(p.RightNode, q.RightNode)
}
