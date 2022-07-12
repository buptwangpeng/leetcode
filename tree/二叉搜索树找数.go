package tree

import "leetcode/utils"

// leetcode 简单 700 二叉搜索树找数

func searchBST(root *utils.TreeNode, val int) *utils.TreeNode {
	if root == nil {
		return nil
	}
	if root.Value == val {
		return root
	}
	if root.Value < val {
		return searchBST(root.RightNode, val)
	}
	if root.Value > val {
		return searchBST(root.LeftNode, val)
	}
	return nil
}
