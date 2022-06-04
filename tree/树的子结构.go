package tree

import "leetcode/utils"

// leetcode 剑指 Offer 26. 树的子结构
func IsSubStructure(A *utils.TreeNode, B *utils.TreeNode) bool {
	// 前序遍历
	if A == nil || B == nil {
		return false
	}
	//遍历A中每个节点，A树中任一节点包含B就能返回true
	left := IsSubStructure(A.LeftNode, B)
	right := IsSubStructure(A.RightNode, B)
	return isContain(A, B) || left || right
}

// 这个方法挺关键
//包含：以A为根的树是否包含B
func isContain(A *utils.TreeNode, B *utils.TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Value != B.Value {
		return false
	}
	return isContain(A.LeftNode, B.LeftNode) && isContain(A.RightNode, B.RightNode)
}
