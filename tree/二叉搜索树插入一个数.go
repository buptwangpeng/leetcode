package tree

import "leetcode/utils"

// leetcode 中等 701 二叉搜索树插入一个数
func insertIntoBST(root *utils.TreeNode, val int) *utils.TreeNode {
	// 背模版，这道题可以背会，方法感觉不能完全理解
	// 找到空位置，插入新节点
	if root == nil {
		return &utils.TreeNode{Value: val}
	}
	if root.Value == val {
		return root
	}
	// val大，应插到右子树上
	if root.Value < val {
		root.RightNode = insertIntoBST(root.RightNode, val)
	}
	// val小，应插到左子树上
	if root.Value > val {
		root.LeftNode = insertIntoBST(root.LeftNode, val)
	}
	return root

}
