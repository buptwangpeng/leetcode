package tree

import "leetcode/utils"

// leetcode 剑指 Offer 27. 二叉树的镜像

func MirrorTree(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}
	root.LeftNode, root.RightNode = root.RightNode, root.LeftNode
	MirrorTree(root.LeftNode)
	MirrorTree(root.RightNode)
	return root
}
