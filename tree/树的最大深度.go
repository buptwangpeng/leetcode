package tree

import "leetcode/utils"

// leetcode 104 easy 树的最大深度

func maxDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.LeftNode)
	right := maxDepth(root.RightNode)
	return max(left, right) + 1
}
