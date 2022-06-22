package tree

import "leetcode/utils"

// leetcode 中等 98

// 核心思路：相当于给树上的所有节点添加了一个min和max边界，约束root的左子树节点值小于root的值，右子树的值大于root的值
func isValidBST(root *utils.TreeNode) bool {
	return isValidBST0(root, nil, nil)
}
func isValidBST0(root, min, max *utils.TreeNode) bool {
	// 处理root
	if root == nil {
		return true
	}
	if min != nil && root.Value <= min.Value {
		return false
	}
	if max != nil && root.Value >= max.Value {
		return false
	}

	// 递归
	return isValidBST0(root.LeftNode, min, root) && isValidBST0(root.RightNode, root, max)
}
