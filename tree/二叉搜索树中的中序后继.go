package tree

import "leetcode/utils"

// leetcode  剑指 Offer II 053. 二叉搜索树中的中序后继

var res0 *utils.TreeNode
var ok0 bool

func InorderSuccessor(root *utils.TreeNode, p *utils.TreeNode) *utils.TreeNode {
	res0 = nil
	ok0 = false
	inorderSuccessor0(root, p)
	return res0
}

func inorderSuccessor0(root *utils.TreeNode, p *utils.TreeNode) {
	// 中序遍历
	if root == nil {
		return
	}
	inorderSuccessor0(root.LeftNode, p)
	if root.Value == p.Value {
		ok0 = true
	}
	// fmt.Println("root.Val: ", root.Val)
	if root.Value != p.Value && ok0 {
		// fmt.Println("suceess: ", root.Val)
		res0 = root
		ok0 = false
	}
	inorderSuccessor0(root.RightNode, p)
}
