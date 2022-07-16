package tree

import (
	"leetcode/utils"
)

// leetcode 226. 翻转二叉树 简单

// 先背写树的模版
func invertTree(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}
	root.LeftNode, root.RightNode = root.RightNode, root.LeftNode
	invertTree(root.RightNode)
	invertTree(root.LeftNode)
	return root
}
