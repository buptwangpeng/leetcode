package tree

import "leetcode/utils"

// leetcode 99 中等
var prev, x, y *utils.TreeNode

func RecoverTree(root *utils.TreeNode) {
	prev, x, y = nil, nil, nil
	recoverTree0(root)
	x.Value, y.Value = y.Value, x.Value
}
func recoverTree0(root *utils.TreeNode) {
	if root == nil {
		return
	}
	recoverTree0(root.LeftNode)

	// 中序遍历
	if prev != nil && root.Value < prev.Value {
		y = root
		if x == nil {
			x = prev
		}
	}
	prev = root

	recoverTree0(root.RightNode)
}
